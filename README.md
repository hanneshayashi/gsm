## GSM

GoSpace Manager - Manage Google Workspace resources using a simple CLI written in Go.

### Introduction

I created this tool first and foremost for myself for the following reasons:
- I wanted to get better at Go
- The existing tools like [GAM](https://github.com/jay0lee/GAM) and [PSGSuite](https://github.com/SCRT-HQ/PSGSuite), while excellent, did not fit my use cases sometimes and I wanted to have the flexibility to quickly make changes myself
- Programming is fun!

If you need a mature and battle-tested tool with an active community, feel free to try one of the linked alternatives.\
GSM does exactly what **I** need, but that does not necessarily mean that it will be useful to everyone.\
However, by making it open source, I hope that I can help someone looking for a solution that other tools don't offer.

### General Information

GSM was intentionally designed to be as close to the actual APIs as possible.\
Because of that, it may not be as user-friendly as some of the alternatives out there.\
On the plus-side, you can usually look at the Google API documentation that is linked to in every command's description to figure out how something works. Most flag descriptions are also taken from the official API docs, so I take no credit there!

### License and Copyright

GoSpace Manager (GSM) is licensed under the [GPLv3](https://gsm.hayashi-ke.online/license) as free software.
Copyright © 2020 Hannes Hayashi.

### Third Party Libraries

GSM is based on open source technology and would not exist without the incredible work of some people:
- The engineers at Google who created the APIs and Go libraries GSM is based on
  - https://github.com/googleapis/google-api-go-client
- GSM uses Cobra and Viper for command and configuration management
  - https://github.com/spf13/cobra
- Even though there are many alternatives, I found the simple and elegant library by flowchartsman for exponential backoff and retry to be just perfect (it even supports concurrency!)
  - https://github.com/flowchartsman/retry

### Features

GSM offers CLI access to over 20 Google API via 50+ main commands and over 
250 sub commands (plus some 230 custom "batch" commands).

You can use GSM in one of two modes
- user: User mode allows you to use any Google account (even private ones) to access the APIs.\
        Note that you will only have access to the resources and APIs your account can access!
- dwd:  DWD (Domain Wide Delegation) allows you to utilize a GCP service account to impersonate user accounts in a G Suite domain.\
 You need to add the service account and the appropriate scopes in the Admin Console of your G Suite domain to us this mode.

You can set up multiple configurations using [gsm configs](https://gsm.hayashi-ke.online/gsm/configs).

### Output

GSM is a CLI for the official Google API. It is designed to be easily usable in scripts and workflows. To that end, I made the decision to ommit the implementation of "interesting" output that tells you what GSM is doing, because, while it may be neat to watch, it doesn't serve a purpose when you want to create a script that actually uses the output and I hate the idea of parsing unformatted text to make decisions. Therefore, all* of GSM's console output is parseable JSON or XML (mostly what the API returns).

*the [configs](https://gsm.hayashi-ke.online/gsm/configs) command is a notable exception.

#### Scripting examples

I highly recommend considering using GSM together with PowerShell or Python when creating scripts.\
GSM works nicely with PowerShell's ConvertFrom-Json commandlet.

You can take a look at some examples under [scripting](https://gsm.hayashi-ke.online/scripting).

#### Logging

As useful as the above may be, sometimes you need to understand what is happening or need to know why something didn't work as expected. For those time, GSM creates a log file in the same directory as the binary called "gsm.log" that contains error messages.

### See Also

* [setup](https://gsm.hayashi-ke.online/setup)       - How to set up GSM
* [gsm](https://gsm.hayashi-ke.online/gsm)   - Command overview
* [batch commands](https://gsm.hayashi-ke.online/batch_commands)     - How to use batch commands
* [examples](https://gsm.hayashi-ke.online/examples) - See some examples
* [scripting examples](https://gsm.hayashi-ke.online/scripting) - Some examples on how to use GSM in scripts