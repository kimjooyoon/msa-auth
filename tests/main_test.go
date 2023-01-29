package tests

import (
	"bytes"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"log"
	"msa-auth/api"
	"msa-auth/database"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func mariadbSetUp(ctx context.Context) (testcontainers.Container, int, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mariadb:latest",
		Env:          map[string]string{"MARIADB_ROOT_PASSWORD": "test"},
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("mariadbd: ready for connections."),
	}
	mariaC, err1 := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err1 != nil {
		return nil, 0, err1
	}
	port, err2 := mariaC.MappedPort(ctx, "3306/tcp")
	if err2 != nil {
		return nil, 0, err2
	}

	return mariaC, port.Int(), nil
}

func redisSetUp(ctx context.Context) (testcontainers.Container, int, error) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err1 := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err1 != nil {
		return nil, 0, err1
	}
	port, err2 := redisC.MappedPort(ctx, "6379/tcp")
	if err2 != nil {
		return nil, 0, err2
	}

	return redisC, port.Int(), nil
}

func TestMain(m *testing.M) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panicf("Error loading .env file\n%v", errEnv)
	}

	err := os.Setenv("PORT", "49999")
	if err != nil {
		log.Panicf("os.Setenv err\n%v", err)
	}
	err = os.Setenv("HOST", "localhost")
	if err != nil {
		log.Panicf("os.Setenv err\n%v", err)
	}
	ctx := context.Background()

	mariaC, port1, err1 := mariadbSetUp(ctx)
	if err1 != nil {
		log.Panicf("mariadbSetUp(ctx) err\n%v", err1)
	}

	redisC, port2, err2 := redisSetUp(ctx)
	if err2 != nil {
		log.Panicf("redisSetUp(ctx) err\n%v", err2)
	}
	err2 = os.Setenv("REDIS_DSN", "localhost:"+strconv.Itoa(port2))
	if err != nil {
		log.Panicf("os.Setenv(\"REDIS_DSN\", \"localhost:\"+strconv.Itoa(port2)) err\n%v", err2)
	}

	if err2 != nil {
		log.Panicf("%v", err2)
	}

	defer func() {
		if err := mariaC.Terminate(ctx); err != nil {
			log.Panicf("%v", err)
		}
	}()
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Panicf("%v", err)
		}
	}()
	time.Sleep(3 * time.Second)
	tmp := database.MysqlConnection(database.DSN("root:test@tcp(localhost:" + strconv.Itoa(port1) + ")/sys"))
	tmp.Exec("create schema msa_auth")
	database.Clear()

	err = os.Setenv("DSN", "root:test@tcp(localhost:"+strconv.Itoa(port1)+
		")/msa_auth?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panicf("os.Setenv(\"DSN\", \"root:test@tcp(localhost:\"+strconv.Itoa(port1)+\n\t\t\")/msa_auth?charset=utf8mb4&parseTime=True&loc=Local\")\n%v", err)
	}
	database.AutoMigrate()
	database.Clear()

	go api.RunServer(false, "*")
	time.Sleep(3 * time.Second)

	code := m.Run()
	os.Exit(code)
}

func getRequest(path, method string, b []byte) (*http.Client, *http.Request) {
	reqBody := bytes.NewReader(b)
	url := fmt.Sprintf("http://%s:%s%s", os.Getenv("HOST"), os.Getenv("PORT"), path)
	c := http.DefaultClient
	req, err1 := http.NewRequest(method, url, reqBody)
	if err1 != nil {
		log.Panicf("%v", err1)
	}
	return c, req
}

func ClientE(path, method string, b []byte) (string, string) {
	c, req := getRequest(path, method, b)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicf("io.ReadCloser error %v", err)
		}
	}(resp.Body)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	expected := string(respBody)
	return expected, resp.Status
}

func ClientToken(path, method string, b []byte, token string) (string, string) {
	c, req := getRequest(path, method, b)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicf("io.ReadCloser error %v", err)
		}
	}(resp.Body)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	expected := string(respBody)
	return expected, resp.Status
}
