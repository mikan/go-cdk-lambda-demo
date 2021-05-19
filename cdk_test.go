package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestGoCDKDemoFunctionStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewGoCDKDemoFunctionStack(app, "MyStack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	displayName := template.Get("Resources.GoCDKDemoFunctionBD53E519.Properties.FunctionName").String()
	assert.Equal(t, "GoCDKDemoFunction", displayName)
}
