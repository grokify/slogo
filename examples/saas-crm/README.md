# SaaS CRM Business Journey Metrics Examples

This directory contains OpenSLO examples for SaaS CRM platforms (like Salesforce or HubSpot), tracking the complete user journey from acquisition through retention and measuring business outcomes.

## What are SaaS CRM Journey Metrics?

These metrics track the end-to-end user experience and business performance of a CRM platform, measuring how users progress through their journey, engage with features, and deliver business value.

## Files in this Directory

### user-activation.go
New user onboarding and activation:
- `ExampleUserActivationSLO` - Percentage of users completing activation within 7 days
- `ExampleTimeToFirstValueSLO` - How quickly users create their first contact/deal
- `ExampleOnboardingCompletionSLO` - Onboarding checklist completion rate

### user-engagement.go
Core engagement metrics (MAU, DAU, stickiness):
- `ExampleDailyActiveUsersSLO` - **DAU** (Daily Active Users)
- `ExampleMonthlyActiveUsersSLO` - **MAU** (Monthly Active Users)
- `ExampleDAUMAURatioSLO` - **DAU/MAU ratio** (stickiness metric)
- `ExampleWeeklyActiveUsersSLO` - **WAU** (Weekly Active Users)
- `ExampleDAUWAURatioSLO` - **DAU/WAU ratio** (weekly stickiness)
- `ExamplePowerUserRatioSLO` - Percentage of highly engaged users

### feature-adoption.go
Feature usage and adoption tracking:
- `ExampleContactManagementUsageSLO` - Contact feature adoption
- `ExampleDealPipelineUsageSLO` - Deal pipeline feature adoption
- `ExampleEmailIntegrationUsageSLO` - Email integration adoption
- `ExampleReportingUsageSLO` - Analytics/reporting feature usage
- `ExampleMobileAppUsageSLO` - Mobile app adoption rate

### business-outcomes.go
Business value and productivity metrics:
- `ExampleDealsCreatedSLO` - Deal creation velocity
- `ExampleDealWinRateSLO` - Percentage of deals won
- `ExampleSalesCycleLengthSLO` - Average time to close deals
- `ExampleContactCreationRateSLO` - Contact growth rate
- `ExampleEmailsSentPerUserSLO` - User productivity metric

### user-retention.go
Retention, churn, and cohort analysis:
- `ExampleDay7RetentionSLO` - 7-day user retention
- `ExampleDay30RetentionSLO` - 30-day user retention
- `ExampleChurnRateSLO` - Monthly customer churn rate
- `ExampleResurrectionRateSLO` - Dormant user reactivation
- `ExampleCohortRetentionSLO` - 90-day cohort retention

### saas_crm_test.go
Validation tests for all 25 SaaS CRM SLOs

## The SaaS Customer Journey

### 1. Acquisition → Activation
**Goal**: Get users to "aha moment" quickly
- Track: Activation rate, time to first value, onboarding completion
- Target: 50-70% activation within 7 days

### 2. Activation → Engagement
**Goal**: Build regular usage habits
- Track: DAU, WAU, MAU, DAU/MAU ratio
- Target: 30-50% DAU/MAU ratio (good stickiness)

### 3. Engagement → Adoption
**Goal**: Drive feature discovery and multi-feature usage
- Track: Feature adoption rates across core capabilities
- Target: 60-80% adoption of core features

### 4. Adoption → Value
**Goal**: Deliver measurable business outcomes
- Track: Deals created, win rates, productivity metrics
- Target: Increasing trend in business outcomes

### 5. Value → Retention
**Goal**: Prevent churn, maximize LTV
- Track: Day 7/30 retention, churn rate, cohort retention
- Target: >40% day-30 retention, <3% monthly churn

## Key Metric Definitions

### Engagement Metrics

**DAU (Daily Active Users)**
- Users who perform any action in the platform on a given day
- Indicates daily engagement level

**MAU (Monthly Active Users)**
- Users who perform any action in the platform in a 30-day period
- Indicates overall platform reach

**DAU/MAU Ratio (Stickiness)**
- DAU ÷ MAU = percentage of monthly users who are active daily
- Higher ratio = more engaged, "sticky" product
- Benchmarks:
  - 10-20% = Low stickiness (monthly tools)
  - 20-40% = Medium stickiness (weekly tools)
  - 40%+ = High stickiness (daily tools)

**WAU (Weekly Active Users)**
- Users active in a 7-day period
- Useful mid-range metric between DAU and MAU

**Power Users**
- Users active 20+ days per month
- Typically 15-30% of MAU for healthy products

### Retention Metrics

**Day N Retention**
- Percentage of users active on day N after signup
- Critical milestones: Day 1, 7, 30, 90

**Cohort Retention**
- Track retention by signup cohort over time
- Reveals product improvements across cohorts

**Churn Rate**
- Percentage of customers lost in a period
- SaaS benchmark: 3-7% monthly churn (B2B)

## Example Queries

### Prometheus Queries
```promql
# DAU
count(count_over_time(user_activity_total[24h]) > 0)

# MAU
count(count_over_time(user_activity_total[30d]) > 0)

# DAU/MAU Ratio
count(count_over_time(user_activity_total[24h]) > 0) / count(count_over_time(user_activity_total[30d]) > 0)
```

### BigQuery Queries
```sql
-- Day 7 Retention
SELECT
  COUNT(DISTINCT user_id)
FROM user_retention
WHERE days_since_signup = 7
  AND was_active = true

-- Power User Ratio
SELECT
  COUNT(DISTINCT user_id)
FROM user_activity
WHERE days_active_last_30d >= 20
```

## When to Use These Metrics

SaaS CRM journey metrics are essential for:
- **CRM platforms**: Salesforce, HubSpot, Pipedrive
- **Sales tools**: Outreach, SalesLoft, Apollo
- **Customer success platforms**: Gainsight, ChurnZero
- **Marketing automation**: Marketo, Pardot
- **Any B2B SaaS** with similar user journeys

## Key Performance Indicators by Stage

### Early Stage (Product-Market Fit)
- Focus: Activation, Day 7/30 retention
- Target: >50% activation, >40% Day 30 retention

### Growth Stage (Scaling)
- Focus: DAU/MAU, feature adoption, business outcomes
- Target: >30% DAU/MAU, 70%+ core feature adoption

### Mature Stage (Optimization)
- Focus: Churn reduction, expansion revenue, power users
- Target: <3% churn, 25%+ power users

## Common Pitfalls

1. **Vanity metrics**: Don't just track MAU growth without engagement
2. **Missing per-user view**: Aggregate metrics can hide poor individual experiences
3. **Ignoring feature adoption**: Users who adopt 3+ features have 2x retention
4. **Late-stage focus**: Fix activation before optimizing retention

## References

- [Measuring Product Health](https://www.reforge.com/blog/retention-engagement-growth-silent-killer) - Reforge
- [SaaS Metrics 2.0](https://www.forentrepreneurs.com/saas-metrics-2/) - For Entrepreneurs
- [Product-Led Growth Benchmarks](https://openviewpartners.com/product-led-growth/) - OpenView
- [The Lean Analytics Book](https://leananalyticsbook.com/) - Alistair Croll & Benjamin Yoskovitz
- OpenSLO Specification: https://github.com/OpenSLO/OpenSLO
