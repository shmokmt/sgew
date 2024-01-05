# sgew

SendGrid Event Webhook Debugger

## Requirements

```
export SENDGRID_API_KEY=xxx
export NGROK_AUTHTOKEN=xxx
```

## Usage

```
❯ sgew --help
Usage: sgew <command>

SendGrid Event Webhook Debugger

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  listen
    Listen for webhook events

  trigger --id=STRING --url=STRING
    trigger test webhook events

  ls
    List all event webhooks

Run "sgew <command> --help" for more information on a command.
```

List all event webhooks

```
+--------------------------------------+---------+------------------------------------------------------------------------------------------------------------------+------------------------------------------+--------------+
| ID                                   | ENABLED | URL                                                                                                              | FRIENDLY NAME                            | CREATED DATE |
+--------------------------------------+---------+------------------------------------------------------------------------------------------------------------------+------------------------------------------+--------------+
| 61d7a811-1855-41bc-a75c-d5ef9d1adefa | false   | https://receiver1.example.com                                                                                    | receiver1                                | 2023-10-24   |
+--------------------------------------+---------+------------------------------------------------------------------------------------------------------------------+------------------------------------------+--------------+
| b2b55c5d-d703-494c-b6c3-4618bfb5dee9 | true    | https://receiver2.example.com                                                                                    | receiver2                                | 2023-12-19   |
+--------------------------------------+---------+------------------------------------------------------------------------------------------------------------------+------------------------------------------+--------------+
```

# See Also

- https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/authentication
- https://ngrok.com/docs/agent/cli/
