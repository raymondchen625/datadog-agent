// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !windows

package proc

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPidListInvalid(t *testing.T) {
	pids := getPidList("/incorrect/folder")
	assert.Equal(t, 0, len(pids))
}

func TestGetPidListValid(t *testing.T) {
	pids := getPidList("./testData")
	sort.Ints(pids)
	assert.Equal(t, 2, len(pids))
	assert.Equal(t, 13, pids[0])
	assert.Equal(t, 142, pids[1])
}

func TestSearchProcsForEnvVariableFromPidIncorrect(t *testing.T) {
	envVars := getEnvVariablesFromPid("./testData", 999)
	assert.Equal(t, 0, len(envVars))
}

func TestSearchProcsForEnvVariableFromPidCorrect(t *testing.T) {
	envVars := getEnvVariablesFromPid("./testData", 13)
	assert.Equal(t, "value0", envVars["env0"])
	assert.Equal(t, "value1", envVars["env1"])
	assert.Equal(t, "AWS_Lambda_nodejs14.x", envVars["AWS_EXECUTION_ENV"])
	assert.Equal(t, "value3", envVars["env3"])
	assert.Equal(t, 4, len(envVars))
}

func TestSearchProcsForEnvVariableFound(t *testing.T) {
	result := SearchProcsForEnvVariable("./testData", "env1")
	expected := []string{"value1"}
	assert.Equal(t, 1, len(result))
	assert.Equal(t, expected[0], result[0])
}
func TestSearchProcsForEnvVariableNotFound(t *testing.T) {
	result := SearchProcsForEnvVariable("./testData", "xxx")
	assert.Equal(t, 0, len(result))
}

func TestParseCPUTotals(t *testing.T) {
	path := "./testData/valid_stat"
	userCPUTimeMs, systemCPUTimeMs, err := getCPUData(path)
	assert.Equal(t, float64(23370), userCPUTimeMs)
	assert.Equal(t, float64(1880), systemCPUTimeMs)
	assert.Nil(t, err)

	path = "./testData/invalid_stat_non_numerical_value_1"
	userCPUTimeMs, systemCPUTimeMs, err = getCPUData(path)
	assert.Equal(t, float64(0), userCPUTimeMs)
	assert.Equal(t, float64(0), systemCPUTimeMs)
	assert.NotNil(t, err)

	path = "./testData/invalid_stat_non_numerical_value_2"
	userCPUTimeMs, systemCPUTimeMs, err = getCPUData(path)
	assert.Equal(t, float64(0), userCPUTimeMs)
	assert.Equal(t, float64(0), systemCPUTimeMs)
	assert.NotNil(t, err)

	path = "./testData/invalid_stat_wrong_number_columns"
	userCPUTimeMs, systemCPUTimeMs, err = getCPUData(path)
	assert.Equal(t, float64(0), userCPUTimeMs)
	assert.Equal(t, float64(0), systemCPUTimeMs)
	assert.NotNil(t, err)

	path = "./testData/nonexistant_stat"
	userCPUTimeMs, systemCPUTimeMs, err = getCPUData(path)
	assert.Equal(t, float64(0), userCPUTimeMs)
	assert.Equal(t, float64(0), systemCPUTimeMs)
	assert.NotNil(t, err)
}

func TestGetNetworkData(t *testing.T) {
	path := "./testData/net/valid_dev"
	networkData, err := getNetworkData(path)
	assert.Nil(t, err)
	assert.Equal(t, float64(180), networkData.RxBytes)
	assert.Equal(t, float64(254), networkData.TxBytes)

	path = "./testData/net/invalid_dev_malformed"
	networkData, err = getNetworkData(path)
	assert.NotNil(t, err)
	assert.Nil(t, networkData)

	path = "./testData/net/invalid_dev_non_numerical_value"
	networkData, err = getNetworkData(path)
	assert.NotNil(t, err)
	assert.Nil(t, networkData)

	path = "./testData/net/missing_interface_dev"
	networkData, err = getNetworkData(path)
	assert.NotNil(t, err)
	assert.Nil(t, networkData)

	path = "./testData/net/nonexistent_dev"
	networkData, err = getNetworkData(path)
	assert.NotNil(t, err)
	assert.Nil(t, networkData)
}
