# AI Agents

Comprehensive monitoring for AI agent platforms, covering both end-user experience metrics and business performance indicators.

## What are AI Agent Metrics?

AI agent metrics track the health, quality, and business impact of AI-powered agents. These metrics go beyond traditional service metrics to include AI-specific concerns like:

- Response quality and accuracy
- Hallucination detection
- Task completion success
- Cost efficiency (token usage)
- Per-user fairness

## SLOs in this Example (20 total)

### Availability & Reliability (2 SLOs)

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleAgentAvailabilitySLO` | Aggregated | Overall service availability |
| `ExamplePerUserAgentAvailabilitySLO` | Per-user | Consistent per-user availability |

### Response Quality (3 SLOs)

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleAgentResponseQualitySLO` | Aggregated | User satisfaction ratings |
| `ExamplePerUserResponseQualitySLO` | Per-user | Per-user quality tracking |
| `ExampleAgentAccuracySLO` | Accuracy | Hallucination rate |

### Performance (3 SLOs)

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleAgentResponseTimeSLO` | Aggregated | P95 response time |
| `ExamplePerUserResponseTimeSLO` | Per-user | Per-user latency |
| `ExampleAgentFirstTokenLatencySLO` | Streaming | Time to first token |

### Task Completion (4 SLOs)

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleTaskCompletionRateSLO` | Aggregated | Task success rate |
| `ExamplePerUserTaskCompletionSLO` | Per-user | Per-user task success |
| `ExampleTaskAbandonmentRateSLO` | Aggregated | User frustration indicator |
| `ExampleMultiStepTaskSuccessSLO` | Complex | Multi-step task completion |

### User Engagement (4 SLOs)

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleDailyActiveUsersSLO` | DAU | Daily active users |
| `ExampleUserRetentionSLO` | Retention | 7-day retention rate |
| `ExampleSessionDurationSLO` | Engagement | Average session length |
| `ExampleConversationTurnsSLO` | Depth | Conversation depth |

### Cost Efficiency (4 SLOs)

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleTokenUsagePerTaskSLO` | Efficiency | Tokens per task |
| `ExamplePerUserCostSLO` | Per-user | Per-user cost tracking |
| `ExampleCostPerSuccessfulTaskSLO` | ROI | Cost per successful task |
| `ExampleCacheHitRateSLO` | Caching | Response cache efficiency |

## Usage

```go
import aiagents "github.com/grokify/slogo/examples/ai-agents"

// Get individual SLOs
qualitySLO := aiagents.ExampleAgentResponseQualitySLO()
costSLO := aiagents.ExampleTokenUsagePerTaskSLO()

// Get all SLOs
slos := aiagents.SLOs()
```

## Key Metric Categories

### 1. Availability & Reliability

Monitor whether agents are accessible and functional when users need them.

!!! note "Per-User Fairness"
    Per-user SLOs ensure no single user experiences consistently poor service, even if aggregate metrics look healthy.

### 2. Quality & Accuracy

Track response quality, user satisfaction, and accuracy (hallucination detection).

```promql
# Hallucination rate
1 - (sum(rate(agent_responses_factual_total[24h])) / sum(rate(agent_responses_total[24h])))
```

### 3. Performance

Response times for both batch and streaming responses.

```promql
# Time to first token (streaming)
histogram_quantile(0.95, rate(agent_first_token_latency_seconds_bucket[5m]))
```

### 4. Task Completion

Success rates for user-initiated tasks.

### 5. Cost Efficiency

Token usage, caching, and cost per outcome.

```promql
# Average tokens per task
sum(rate(llm_tokens_used_total[24h])) / sum(rate(agent_tasks_completed_total[24h]))
```

## Ontology Labels

AI Agent SLOs use mixed labels depending on the metric:

```go
// Service-level (availability, performance)
ontology.LabelFramework:  ontology.FrameworkRED,
ontology.LabelLayer:      ontology.LayerService,
ontology.LabelAudience:   ontology.AudienceSRE,

// Business-level (engagement, cost)
ontology.LabelFramework:  ontology.FrameworkCustom,
ontology.LabelLayer:      ontology.LayerBusiness,
ontology.LabelAudience:   ontology.AudienceProduct,
business.LabelDomain:     business.DomainAIML,
```

## Per-User vs Aggregated

| Metric Type | Aggregated | Per-User |
|-------------|------------|----------|
| Availability | 99.9% overall | 99% per-user minimum |
| Response Time | P95 < 2s | 95% of users < 3s |
| Quality | 4.5/5 average | 4.0/5 per-user minimum |
| Task Success | 85% overall | 80% per-user minimum |

## References

- 🤖 [AI Agent Monitoring Best Practices](https://langchain.com/docs/best-practices/monitoring)
- 📊 [LLM Observability](https://www.anthropic.com/engineering/measuring-model-quality)
- 📜 [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)
