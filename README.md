# github-billing-exporter
github-billing-exporter for prometheus
Referance: [https://docs.github.com/en/free-pro-team@latest/rest/reference/billing](https://docs.github.com/en/free-pro-team@latest/rest/reference/billing)

## Options
| Name | Flag | Env vars | Default | Description |
|---|---|---|---|---|
| Github Token | token, t | TOKEN | - | Personnal Access Token. Organization mode must have the `repo` or `admin:org` scope, User mode must have the `user` scope. |
| Github Organization | organization, o | ORGANIZATION | - | Organization name to get GitHub billing report, mutually exclusive with User |
| Github User | user, u | USER | - | User name to get GitHub billing report, mutually exclusive with Organization |
| Refresh | refresh, r | REFRESH | 300 | Refresh time fetch GitHub billing report in sec |
| Exporter port | port, p | PORT | 9999 | Exporter port |

## Exported stats
### GitHub Actions total_minutes_used
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Minutes | Number of total minutes used during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Actions total_paid_minutes_used
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Minutes | Number of paid minutes used during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Actions included_minutes
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Minutes | Number of included minutes during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Actions minutes_used_breakdown
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Minutes | Number of minutes used breakdown during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |
| os | Runner OS(ubuntu, macos or windows). |

### GitHub Pakcages total_gigabytes_bandwidth_used
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Gigabytes | Number of total gigabytes bandwidth used during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Pakcages total_paid_gigabytes_bandwidth_used
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Gigabytes | Number of total paid gigabytes bandwidth used during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Pakcages included_gigabytes_bandwidth
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Gigabytes | Number of included gigabytes bandwidth during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Shared Storage days_left_in_billing_cycle
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Days | Number of days left during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Shared Storage estimated_paid_storage_for_month
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Gigabytes | Number of estimated paid storage during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

### GitHub Shared Storage estimated_storage_for_month
Gauge type

#### Result possibility
| Gauge | Description |
| --- | --- |
| Gigabytes | Number of estimated storage during the current billing cycle. |

#### Fieldes
| Name | Description |
| --- | --- |
| owner | Billing owner(Organization Name or User Name). |

## Usage
```bash
Starts GitHubBillingExporter as a server

Usage:
  github-billing-exporter server [flags]

Flags:
  -h, --help                  help for server
  -o, --organization string   GitHub Organization Name
  -p, --port int              Exporter Listen Port (default 9999)
  -r, --refresh int           Refresh Interval Secounds (default 300)
  -t, --token string          GitHub Token
  -u, --user string           GitHub User Name
```
