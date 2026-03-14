# Domain Packages

Domain-specific ontology labels are provided as opt-in packages under `ontologies/domains/`. Import only the domains you need.

## Available Domains

| Package | Import Path | Use Case |
|---------|-------------|----------|
| [IAM](#iam) | `ontologies/domains/iam` | Identity & Access Management |
| [Product](#product) | `ontologies/domains/product` | Product metrics, AARRR |
| [Business](#business) | `ontologies/domains/business` | Industry verticals |
| [SaaS](#saas) | `ontologies/domains/saas` | Multi-tenancy |
| [Security](#security) | `ontologies/domains/security` | Risk & threat |
| [Compliance](#compliance) | `ontologies/domains/compliance` | Regulatory frameworks |
| [Operations](#operations) | `ontologies/domains/operations` | SRE, incident response |

---

## IAM

Identity & Access Management labels for IGA (Identity Governance & Administration) systems.

```go
import "github.com/grokify/slogo/ontologies/domains/iam"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelProcess` | IGA process type |
| `LabelAccessModel` | Access control model |
| `LabelRiskLevel` | Risk classification |
| `LabelRiskType` | Type of risk |
| `LabelCapability` | IGA capability |
| `LabelWorkflow` | Workflow type |
| `LabelImpact` | Business impact |

### Process Values

| Value | Description |
|-------|-------------|
| `access-request` | Access request handling |
| `provisioning` | Account/role provisioning |
| `deprovisioning` | Access removal |
| `recertification` | Periodic access review |
| `sod-check` | Separation of duties |
| `password-management` | Password operations |
| `mfa-enrollment` | MFA setup |
| `session-management` | Session handling |

### Access Model Values

| Value | Description |
|-------|-------------|
| `rbac` | Role-Based Access Control |
| `abac` | Attribute-Based Access Control |
| `pbac` | Policy-Based Access Control |
| `just-in-time` | JIT access |
| `zero-trust` | Zero Trust model |
| `least-privilege` | Minimal required access |

---

## Product

Product metrics and AARRR (Pirate Metrics) journey stages.

```go
import "github.com/grokify/slogo/ontologies/domains/product"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelJourneyStage` | AARRR funnel stage |
| `LabelMetricType` | Product metric type |
| `LabelImpact` | Product impact |

### Journey Stage Values (AARRR)

| Value | Description | Example Metrics |
|-------|-------------|-----------------|
| `acquisition` | User discovery | Signups, landing page views |
| `activation` | First value | Onboarding completion, time to first value |
| `engagement` | Active usage | DAU, MAU, feature adoption |
| `retention` | Return usage | Day 7/30 retention, churn |
| `revenue` | Monetization | Conversion, ARPU |
| `referral` | Viral growth | Referral rate, NPS |

### Metric Type Values

| Value | Description |
|-------|-------------|
| `satisfaction` | User satisfaction (CSAT, NPS) |
| `stickiness` | DAU/MAU ratio |
| `retention` | User return rate |
| `activation` | First value achievement |
| `adoption` | Feature adoption rate |
| `churn` | User attrition |
| `nps` | Net Promoter Score |
| `ces` | Customer Effort Score |
| `csat` | Customer Satisfaction |

---

## Business

Business verticals and industry-specific labels.

```go
import "github.com/grokify/slogo/ontologies/domains/business"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelDomain` | Business vertical |
| `LabelImpact` | Business impact |

### Domain Values

| Value | Description |
|-------|-------------|
| `ai-ml` | AI/ML platforms |
| `crm` | Customer Relationship Management |
| `saas` | Software as a Service |
| `ecommerce` | E-commerce |
| `fintech` | Financial technology |
| `healthcare` | Healthcare/medical |
| `edtech` | Education technology |
| `martech` | Marketing technology |
| `devtools` | Developer tools |
| `cybersecurity` | Security products |

### Impact Values

| Value | Description |
|-------|-------------|
| `revenue` | Revenue impact |
| `cost` | Cost impact |
| `market-share` | Market position |
| `brand-reputation` | Brand/reputation |
| `customer-satisfaction` | Customer satisfaction |
| `employee-productivity` | Workforce efficiency |
| `regulatory-compliance` | Compliance status |

---

## SaaS

Multi-tenancy and SaaS-specific labels.

```go
import "github.com/grokify/slogo/ontologies/domains/saas"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelTenancy` | Tenancy model |
| `LabelImpact` | SaaS-specific impact |

### Impact Values

| Value | Description |
|-------|-------------|
| `tenant-isolation` | Tenant data separation |
| `noisy-neighbor` | Resource contention |
| `tenant-onboarding` | New tenant provisioning |
| `tenant-offboarding` | Tenant removal |
| `resource-quota` | Usage limits |
| `billing-accuracy` | Metering accuracy |

---

## Security

Security risk and threat classification.

```go
import "github.com/grokify/slogo/ontologies/domains/security"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelRiskLevel` | Risk severity |
| `LabelRiskType` | Risk category |
| `LabelThreatType` | Threat classification |
| `LabelImpact` | Security impact |

### Risk Level Values

| Value | Description |
|-------|-------------|
| `critical` | Immediate action required |
| `high` | Urgent attention |
| `medium` | Scheduled remediation |
| `low` | Monitor |
| `info` | Informational |

### Threat Type Values

| Value | Description |
|-------|-------------|
| `authentication` | Auth bypass/failure |
| `authorization` | Privilege escalation |
| `injection` | Code/SQL injection |
| `xss` | Cross-site scripting |
| `csrf` | Cross-site request forgery |
| `dos` | Denial of service |
| `data-exposure` | Data leakage |

### Impact Values

| Value | Description |
|-------|-------------|
| `data-breach` | Data compromise |
| `unauthorized-access` | Access violation |
| `privilege-escalation` | Elevated privileges |
| `data-loss` | Data destruction |
| `data-exfiltration` | Data theft |
| `credential-compromise` | Credential exposure |

---

## Compliance

Regulatory framework labels.

```go
import "github.com/grokify/slogo/ontologies/domains/compliance"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelFramework` | Regulatory framework |
| `LabelRequirement` | Specific requirement |
| `LabelImpact` | Compliance impact |

### Framework Values

| Value | Description |
|-------|-------------|
| `gdpr` | EU General Data Protection Regulation |
| `sox` | Sarbanes-Oxley Act |
| `hipaa` | Health Insurance Portability |
| `pci-dss` | Payment Card Industry |
| `soc2` | Service Organization Control 2 |
| `iso27001` | Information Security |
| `nist` | NIST Cybersecurity Framework |
| `fedramp` | Federal Risk Authorization |
| `ccpa` | California Consumer Privacy |

### Impact Values

| Value | Description |
|-------|-------------|
| `audit-finding` | Audit issue |
| `regulatory-fine` | Financial penalty |
| `certification-risk` | Certification jeopardy |
| `data-breach-notification` | Breach reporting |

---

## Operations

SRE and operational process labels.

```go
import "github.com/grokify/slogo/ontologies/domains/operations"
```

### Labels

| Label | Description |
|-------|-------------|
| `LabelProcess` | Operational process |
| `LabelImpact` | Operations impact |

### Process Values

| Value | Description |
|-------|-------------|
| `incident-response` | Incident handling |
| `change-management` | Change control |
| `deployment` | Release deployment |
| `rollback` | Deployment reversal |
| `capacity-planning` | Resource planning |
| `disaster-recovery` | DR procedures |
| `backup-restore` | Data backup/restore |
| `maintenance` | Scheduled maintenance |

### Impact Values

| Value | Description |
|-------|-------------|
| `mttr` | Mean Time To Recovery |
| `mttd` | Mean Time To Detect |
| `mttf` | Mean Time To Failure |
| `alert-fatigue` | Alert noise |
| `toil` | Manual operational work |
| `deployment-frequency` | Release velocity |
| `change-failure-rate` | Failed changes |
| `lead-time` | Change lead time |

---

## Example: Combining Domains

```go
import (
    "github.com/grokify/slogo/ontology"
    "github.com/grokify/slogo/ontologies/domains/business"
    "github.com/grokify/slogo/ontologies/domains/product"
    "github.com/grokify/slogo/ontologies/domains/compliance"
)

labels := ontology.NewLabels(map[string]string{
    // Core ontology
    ontology.LabelFramework:  ontology.FrameworkCustom,
    ontology.LabelLayer:      ontology.LayerBusiness,
    ontology.LabelAudience:   ontology.AudienceProduct,

    // Business domain
    business.LabelDomain:     business.DomainFintech,

    // Product domain
    product.LabelJourneyStage: product.JourneyStageActivation,
    product.LabelMetricType:   product.MetricTypeActivation,

    // Compliance domain
    compliance.LabelFramework: compliance.FrameworkPCIDSS,
})
```
