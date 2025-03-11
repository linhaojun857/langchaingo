package agents

import (
	"context"

	"github.com/linhaojun857/langchaingo/chains"
	"github.com/linhaojun857/langchaingo/schema"
	"github.com/linhaojun857/langchaingo/tools"
)

const (
	_finalAnswerAction = "Final Answer:"
)

// Agent is the interface all agents must implement.
type Agent interface {
	// Plan Given an input and previous steps decide what to do next. Returns
	// either actions or a finish.
	Plan(ctx context.Context, intermediateSteps []schema.AgentStep, inputs map[string]string, opts ...chains.ChainCallOption) ([]schema.AgentAction, *schema.AgentFinish, error) //nolint:lll
	GetInputKeys() []string
	GetOutputKeys() []string
	GetTools() []tools.Tool
}
