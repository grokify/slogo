# SaaS CRM

End-to-end user journey metrics for SaaS CRM platforms (like Salesforce or HubSpot), tracking the complete user journey from acquisition through retention.

## What are SaaS CRM Journey Metrics?

These metrics track the end-to-end user experience and business performance of a CRM platform, measuring how users progress through their journey, engage with features, and deliver business value.

## SLOs in this Example (25 total)

### User Activation (3 SLOs)

| SLO | Stage | Description |
|-----|-------|-------------|
| `ExampleUserActivationSLO` | Activation | % completing activation in 7 days |
| `ExampleTimeToFirstValueSLO` | Activation | Time to first contact/deal |
| `ExampleOnboardingCompletionSLO` | Activation | Onboarding checklist completion |

### User Engagement (6 SLOs)

| SLO | Metric | Description |
|-----|--------|-------------|
| `ExampleDailyActiveUsersSLO` | DAU | Daily active users |
| `ExampleMonthlyActiveUsersSLO` | MAU | Monthly active users |
| `ExampleDAUMAURatioSLO` | Stickiness | DAU/MAU ratio |
| `ExampleWeeklyActiveUsersSLO` | WAU | Weekly active users |
| `ExampleDAUWAURatioSLO` | Weekly Stickiness | DAU/WAU ratio |
| `ExamplePowerUserRatioSLO` | Engagement | Power users (20+ days/month) |

### Feature Adoption (5 SLOs)

| SLO | Feature | Description |
|-----|---------|-------------|
| `ExampleContactManagementUsageSLO` | Contacts | Contact feature adoption |
| `ExampleDealPipelineUsageSLO` | Deals | Deal pipeline adoption |
| `ExampleEmailIntegrationUsageSLO` | Email | Email integration adoption |
| `ExampleReportingUsageSLO` | Reports | Analytics feature usage |
| `ExampleMobileAppUsageSLO` | Mobile | Mobile app adoption |

### Business Outcomes (5 SLOs)

| SLO | Metric | Description |
|-----|--------|-------------|
| `ExampleDealsCreatedSLO` | Velocity | Deal creation rate |
| `ExampleDealWinRateSLO` | Win Rate | Percentage of deals won |
| `ExampleSalesCycleLengthSLO` | Efficiency | Time to close deals |
| `ExampleContactCreationRateSLO` | Growth | Contact growth rate |
| `ExampleEmailsSentPerUserSLO` | Productivity | Emails per user |

### User Retention (6 SLOs)

| SLO | Metric | Description |
|-----|--------|-------------|
| `ExampleDay7RetentionSLO` | D7 | 7-day retention |
| `ExampleDay30RetentionSLO` | D30 | 30-day retention |
| `ExampleChurnRateSLO` | Churn | Monthly churn rate |
| `ExampleResurrectionRateSLO` | Win-back | Dormant user reactivation |
| `ExampleCohortRetentionSLO` | Cohort | 90-day cohort retention |

## Usage

```go
import saascrm "github.com/grokify/slogo/examples/saas-crm"

// Get individual SLOs
dauSLO := saascrm.ExampleDailyActiveUsersSLO()
stickinessSLO := saascrm.ExampleDAUMAURatioSLO()

// Get all SLOs
slos := saascrm.SLOs()
```

## The SaaS Customer Journey (AARRR)

### 1. Acquisition → Activation

**Goal**: Get users to "aha moment" quickly

- Track: Activation rate, time to first value, onboarding completion
- Target: 50-70% activation within 7 days

### 2. Activation → Engagement

**Goal**: Build daily habits

- Track: DAU, MAU, stickiness (DAU/MAU)
- Target: 40%+ DAU/MAU ratio (good stickiness)

### 3. Engagement → Feature Adoption

**Goal**: Deep product usage

- Track: Feature adoption rates, multi-feature users
- Target: 60%+ using core features weekly

### 4. Feature Adoption → Business Outcomes

**Goal**: Drive customer success

- Track: Deal creation, win rates, productivity
- Target: Positive business impact metrics

### 5. Business Outcomes → Retention

**Goal**: Long-term customer retention

- Track: Retention curves, churn, NRR
- Target: 97%+ monthly retention (3% churn)

## Stickiness Metrics

| Metric | Formula | Good Target |
|--------|---------|-------------|
| DAU/MAU | Daily / Monthly | > 40% |
| DAU/WAU | Daily / Weekly | > 60% |
| Power User Ratio | 20+ days active / MAU | > 25% |

## Ontology Labels

SaaS CRM SLOs use business-layer labels:

```go
import (
    "github.com/grokify/slogo/ontology"
    "github.com/grokify/slogo/ontologies/domains/business"
    "github.com/grokify/slogo/ontologies/domains/product"
)

ontology.LabelFramework:    ontology.FrameworkCustom,
ontology.LabelLayer:        ontology.LayerBusiness,
ontology.LabelScope:        ontology.ScopeBusinessOutcome,
business.LabelDomain:       business.DomainCRM,
product.LabelJourneyStage:  product.JourneyStageActivation, // or Engagement, Retention, Revenue
```

## Data Sources

### Prometheus (Real-time)

```promql
# DAU
count(count_over_time(user_activity_total[24h]) > 0)

# DAU/MAU Ratio
count(count_over_time(user_activity_total[24h]) > 0) /
count(count_over_time(user_activity_total[30d]) > 0)
```

### BigQuery (Cohort Analysis)

```sql
-- Day 7 Retention
SELECT COUNT(DISTINCT user_id)
FROM user_retention
WHERE days_since_signup = 7 AND was_active = true
```

## Industry Benchmarks

| Metric | B2B SaaS Benchmark | Target |
|--------|-------------------|--------|
| DAU/MAU | 13-20% | 40%+ (high engagement) |
| Day 7 Retention | 40-50% | 50%+ |
| Day 30 Retention | 25-35% | 40%+ |
| Monthly Churn | 3-5% | < 3% |
| Activation Rate | 40-60% | 60%+ |

## References

- 📈 [SaaS Metrics 2.0](https://www.forentrepreneurs.com/saas-metrics-2/) - David Skok
- 🏴‍☠️ [AARRR Pirate Metrics](https://www.slideshare.net/dmc500hats/startup-metrics-for-pirates-long-version) - Dave McClure
- 📊 [Measuring Product Health](https://www.reforge.com/blog/retention-engagement-growth-silent-killer) - Reforge
- 📜 [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)
