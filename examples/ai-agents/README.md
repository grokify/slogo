# AI Agents Business Metrics Examples

This directory contains OpenSLO examples for AI agent platforms, covering both end-user experience metrics and business performance indicators.

## What are AI Agent Metrics?

AI agent metrics track the health, quality, and business impact of AI-powered agents that assist users with tasks. These metrics go beyond traditional service metrics to include AI-specific concerns like response quality, accuracy, task completion, and cost efficiency.

## Files in this Directory

### agent-availability.go
Service availability and reliability:
- `ExampleAgentAvailabilitySLO` - Overall agent service availability (aggregated)
- `ExamplePerUserAgentAvailabilitySLO` - Per-user availability to ensure consistent service

### response-quality.go
AI response quality and accuracy:
- `ExampleAgentResponseQualitySLO` - User satisfaction ratings (aggregated)
- `ExamplePerUserResponseQualitySLO` - Per-user quality tracking
- `ExampleAgentAccuracySLO` - Hallucination rate and factual correctness

### response-time.go
Agent performance and latency:
- `ExampleAgentResponseTimeSLO` - P95 response time (aggregated)
- `ExamplePerUserResponseTimeSLO` - Per-user response time consistency
- `ExampleAgentFirstTokenLatencySLO` - Time to first token for streaming responses

### task-completion.go
Task success and completion metrics:
- `ExampleTaskCompletionRateSLO` - Overall task completion rate
- `ExamplePerUserTaskCompletionSLO` - Per-user task success tracking
- `ExampleTaskAbandonmentRateSLO` - User frustration indicator
- `ExampleMultiStepTaskSuccessSLO` - Complex task completion tracking

### user-engagement.go
User activity and retention:
- `ExampleDailyActiveUsersSLO` - DAU tracking
- `ExampleUserRetentionSLO` - 7-day retention rate
- `ExampleSessionDurationSLO` - Average session length
- `ExampleConversationTurnsSLO` - Conversation depth and engagement

### cost-efficiency.go
AI operational costs and efficiency:
- `ExampleTokenUsagePerTaskSLO` - Token efficiency per task
- `ExamplePerUserCostSLO` - Per-user cost tracking
- `ExampleCostPerSuccessfulTaskSLO` - ROI measurement
- `ExampleCacheHitRateSLO` - Response caching efficiency

### ai_agents_test.go
Validation tests for all 20 AI agent SLOs

## Key Metric Categories

### 1. Availability & Reliability
Monitor whether agents are accessible and functional when users need them.

### 2. Quality & Accuracy
Track response quality, user satisfaction, and accuracy (hallucination detection).

### 3. Performance
Measure response times, first-token latency, and overall system responsiveness.

### 4. Task Success
Monitor task completion rates, abandonment, and multi-step task success.

### 5. User Engagement
Track DAU, retention, session depth, and conversation metrics.

### 6. Cost Efficiency
Monitor token usage, per-user costs, and cache hit rates for sustainability.

## Aggregated vs Per-User Metrics

Many metrics include both versions:

**Aggregated Metrics**: Platform-wide view
- Total success rate
- Average response time
- Overall DAU

**Per-User Metrics**: Individual experience view
- Percentage of users with >99% success rate
- Percentage of users with acceptable latency
- Distribution of costs per user

This dual approach helps identify:
- Overall platform health
- Users with poor experiences
- Outliers requiring intervention

## Example Metric Sources

The examples use various data sources:

**Prometheus** (real-time metrics):
```promql
sum(rate(agent_session_starts_total{status="success"}[5m]))
histogram_quantile(0.95, rate(agent_response_duration_seconds_bucket[5m]))
```

**BigQuery** (historical/analytical):
```sql
SELECT COUNT(DISTINCT user_id) FROM returning_users WHERE days_since_first_use <= 7
```

## When to Use These Metrics

AI agent metrics are essential for:
- **AI chatbots and assistants**: Customer support, internal tools
- **Agent platforms**: Multi-agent orchestration systems
- **Workflow automation**: AI-powered task execution
- **Copilot features**: Code assistants, writing assistants
- **Autonomous agents**: Research agents, data analysis agents

## Critical Success Factors for AI Agents

1. **Availability**: Agents must be reliably accessible
2. **Quality**: Responses must be accurate and helpful
3. **Performance**: Fast responses maintain user engagement
4. **Task Success**: Users must accomplish their goals
5. **Cost Control**: LLM costs must be sustainable
6. **User Retention**: Users must find ongoing value

## References

- [Anthropic's AI Safety Best Practices](https://www.anthropic.com/safety)
- [OpenAI's Best Practices for Production](https://platform.openai.com/docs/guides/production-best-practices)
- [Measuring AI Agent Performance](https://www.honeycomb.io/blog/observability-ai-agents)
- OpenSLO Specification: https://github.com/OpenSLO/OpenSLO
