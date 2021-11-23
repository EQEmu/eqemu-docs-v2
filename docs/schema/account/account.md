# account

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBhY2NvdW50X2ZsYWdzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcF9hY2NpZFxuICAgIH1cbiAgICBhY2NvdW50X3Jld2FyZHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50X2lkXG4gICAgfVxuICAgIHNoYXJlZGJhbmsge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2N0aWRcbiAgICB9XG4gICAgYnVnX3JlcG9ydHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50X2lkXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICBpbnQgYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB8fC0tb3sgYWNjb3VudF9mbGFncyA6IEhhcy1NYW55XG4gICAgYWNjb3VudCB8fC0tb3sgYWNjb3VudF9yZXdhcmRzIDogSGFzLU1hbnlcbiAgICBhY2NvdW50IHx8LS1veyBzaGFyZWRiYW5rIDogSGFzLU1hbnlcbiAgICBhY2NvdW50IHx8LS1veyBidWdfcmVwb3J0cyA6IEhhcy1NYW55XG4gICAgYWNjb3VudCB8fC0tb3sgYWNjb3VudF9pcCA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBhY2NvdW50X2ZsYWdzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcF9hY2NpZFxuICAgIH1cbiAgICBhY2NvdW50X3Jld2FyZHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50X2lkXG4gICAgfVxuICAgIHNoYXJlZGJhbmsge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2N0aWRcbiAgICB9XG4gICAgYnVnX3JlcG9ydHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50X2lkXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICBpbnQgYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB8fC0tb3sgYWNjb3VudF9mbGFncyA6IEhhcy1NYW55XG4gICAgYWNjb3VudCB8fC0tb3sgYWNjb3VudF9yZXdhcmRzIDogSGFzLU1hbnlcbiAgICBhY2NvdW50IHx8LS1veyBzaGFyZWRiYW5rIDogSGFzLU1hbnlcbiAgICBhY2NvdW50IHx8LS1veyBidWdfcmVwb3J0cyA6IEhhcy1NYW55XG4gICAgYWNjb3VudCB8fC0tb3sgYWNjb3VudF9pcCA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [account_flags](../../schema/account/account_flags.md) | p_accid |
| Has-Many | id | [account_rewards](../../schema/account/account_rewards.md) | account_id |
| Has-Many | id | [sharedbank](../../schema/account/sharedbank.md) | acctid |
| Has-Many | id | [bug_reports](../../schema/admin/bug_reports.md) | account_id |
| Has-Many | id | [account_ip](../../schema/account/account_ip.md) | accid |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Account Identifier |
| name | varchar | Name |
| charname | varchar | Character name last logged in on this account. |
| sharedplat | int | Platinum in Shared Bank. |
| password | varchar | Private loginserver password. |
| status | int | [Status](../../../../server/player/status-levels) |
| ls_id | varchar |  |
| lsaccount_id | int | Loginserver Account Identifier |
| gmspeed | tinyint | GM Speed: 0 = Disabled, 1 = Enabled |
| revoked | tinyint | OOC Revoked: 0 = False, 1 = True |
| karma | int | Karma |
| minilogin_ip | varchar | Minilogin IP Address |
| hideme | tinyint | Hide Me: 0 = Disabled, 1 = Enabled |
| rulesflag | tinyint | Rules Flag |
| suspendeduntil | datetime | Time Suspension of the Account ends |
| time_creation | int | Time Creation UNIX Timestamp |
| expansion | tinyint | Expansion |
| ban_reason | text | Ban Reason |
| suspend_reason | text | Suspension Reason |

