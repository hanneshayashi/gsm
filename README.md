# GSM

GoSpace Manager - Manage Google Workspace resources using a developer-friendly CLI written in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/hanneshayashi/gsm)](https://goreportcard.com/report/github.com/hanneshayashi/gsm)

## Documentation

View complete documentation at https://gsm.hayashi-ke.online/.

## Introduction

GSM is like gcloud for Google Workspace. It is a no-dependency, free and open source command line interface (CLI) for managing Google Workspace resources. You don't need any software except for the [GSM binary](https://github.com/hanneshayashi/gsm/releases) itself and it is easy to [set up](https://gsm.hayashi-ke.online/setup).

GSM is extremely powerful and flexible, especially when used in [scripts](https://gsm.hayashi-ke.online/scripting) to implement your own custom logic for your specific use case. Since GSM supports pretty much every API endpoint that Google offers for Workspace, there shouldn't really be anything you can't do with GSM, as long as the API supports it.\
You can also [configure](https://gsm.hayashi-ke.online/gsm/configs) pretty much every aspect of GSM, because everyone likes having choices.

GSM also supports many custom commands for common tasks to make your life easier, such as
- ["move"](https://gsm.hayashi-ke.online/gsm/files/move/recursive/) (or [copy](https://gsm.hayashi-ke.online/gsm/files/copy/recursive/)) a folder to a Shared Drive
- ["sync"](https://gsm.hayashi-ke.online/gsm/members/set/recursive/) an organizational unit (OU) to a group
- [send emails](https://gsm.hayashi-ke.online/gsm/messages/send/) (with attachments) via the Gmail API (without SMTP auth)
- [get the size of a Shared Drive](https://gsm.hayashi-ke.online/gsm/drives/getsize/) or [Drive folder](https://gsm.hayashi-ke.online/gsm/files/count/recursive/)

and many more.

GSM was intentionally designed to be as close to the actual APIs as possible.\
Because of that, it may not be as user-friendly as some of the alternatives out there.\
On the plus-side, you can usually look at the Google API documentation that is linked to in every command's description to figure out how something works. Most flag descriptions are also taken from the official API docs, so I take no credit there.

## Features

GSM currently supports the following APIs:
- [Admin SDK Directory API](https://developers.google.com/admin-sdk/directory)
- [Admin SDK Groups Settings API](https://developers.google.com/admin-sdk/groups-settings/get_started)
- [Admin SDK Enterprise License Manager API](https://developers.google.com/admin-sdk/licensing/reference/rest)
- [Cloud Identity API](https://cloud.google.com/identity/docs/reference/rest)
- [Gmail API](https://developers.google.com/gmail/api/reference/rest)
- [Gmail Postmaster API](https://developers.google.com/gmail/postmaster/reference/rest)
- [Google Calendar API](https://developers.google.com/calendar/v3/reference)
- [Contact Delegation API](https://developers.google.com/admin-sdk/contact-delegation/guides)
- [Domain Shared Contacts API](https://developers.google.com/admin-sdk/domain-shared-contacts)
- [Google People API](https://developers.google.com/people/api/rest)
- [Google Drive API](https://developers.google.com/drive/api/v3/reference)
- [Google Drive Labels API](https://developers.google.com/drive/labels/reference/rest)
- [Google Sheets API](https://developers.google.com/sheets/api/reference/rest) (partly)

Most of these APIs allow you to manage multiple object types with each object type allowing multiple operations.\
Overall, GSM supports over **65 [main commands](https://gsm.hayashi-ke/gsm)**, with each one representing an API with multiple methods and each method implemented as a sub command. This amounts to over **500 commands in total**, including over **200 ["batch" commands](https://gsm.hayashi-ke.online/batch_commands)** that allow you to utilize CSV files to apply updates to multiple objects in a multi-threaded manner and over **30 ["recursive" commands](https://gsm.hayashi-ke.online/recursive_commands)** that allow you to apply updates to multiple users in one command, by specifying one or more organizational unit(s) (OUs) and/or group(s).

You can use GSM in one of three modes
- user: User mode allows you to use any Google account (even private ones) to access the APIs.\
        Note that you will only have access to the resources and APIs your account can access!
- dwd:  DWD (Domain Wide Delegation) allows you to utilize a GCP service account to impersonate user accounts in a Workspace domain.\
 You need to add the service account and the appropriate scopes in the Admin Console of your Workspace domain to us this mode.
- adc:  ADC ("Application Default Credentials") mode works like DWD mode, but it allows you to utilize Application Default Credentials, such as the implicit credentials of a Compute Engine instance's Service Account or the "application-default" credentials of the Google Cloud SDK (gcloud), to impersonate a Service Account. This means you don't have to manage Service Account key files. You can also use this mode to use a Service Account directly for API access without specifying a subject to impersonate. Please note that most APIs will not work this way. Only very few Google Workspace APIs support direct access by a Service Account but the option is there if you want to try!

See [Setup](https://gsm.hayashi-ke.online/setup) on how to set up GSM in these modes.

You can also set up multiple configurations using [gsm configs](https://gsm.hayashi-ke.online/gsm/configs) and switch between them using [gsm configs load](https://gsm.hayashi-ke.online/gsm/configs/load) or by specifying the name of the config with the `--config` flag.

## Output

GSM is a CLI for the official Google API. It is designed to be easily usable in scripts and workflows. By default, GSM will output errors to `stderr` *and* the log file in your home directory (see [Logging](#logging)). You can change this with the `--errorOutput` flag either when creating your config or when running individual commands. Output of commands will always be sent to `stdout`. Most commands output in JSON format by default in the same way the API would respond.

If you want to use GSM's output in scripts, you may want to consider using the `--compressOutput` flag, to keep GSM from unnecessarily "prettying up" the output. Depending on the tools you use and how you want to build your workflow, you may also want to consider using the `--streamOutput` flag, which will cause GSM to **stream** single objects directly to `stdout`. In many cases, this is significantly faster and uses a lot less memory because GSM does not need to wait for a command to return all objects to build the return object in memory. However, not all application may properly work with a stream of objects.

### Scripting examples

I highly recommend considering using GSM together with PowerShell or Python when creating scripts.\
GSM works nicely with PowerShell's ConvertFrom-Json commandlet (although there are some [issues with very large amounts of data](https://gsm.hayashi-ke.online/scripting/#processing-very-large-amounts-of-data-with-powershell)).

You can take a look at some examples under [scripting](https://gsm.hayashi-ke.online/scripting).

You can also try the auto-generated [PowerShell module](https://github.com/hanneshayashi/gsm-powershell).\
Note that this module is created with [Crescendo](https://github.com/PowerShell/Crescendo), which is also still in beta. However, for an auto-generated module, it seems to work reasonably well. The module also automatically utilizes streaming.

### Logging

GSM creates a log file in your home directory called `gsm.log` that contains error messages.\
You can configure the location and name of the log file, either in your config file (see [configs](https://gsm.hayashi-ke.online/gsm/configs)) or by using the `--log` flag when running a command.
You can also use the [log command](https://gsm.hayashi-ke.online/gsm/log) to view or clear the log, without having to manually open it.

## License and Copyright

GoSpace Manager (GSM) is licensed under the [GPLv3](https://gsm.hayashi-ke.online/license) as free software.\
Copyright © 2020-2023 Hannes Hayashi.

## Third Party Libraries

GSM is based on open source technology and would not exist without the incredible work of some people:
- The engineers at Google who created the APIs and Go libraries GSM is based on
  - https://github.com/googleapis/google-api-go-client
- GSM uses Cobra and Viper for command and configuration management
  - https://github.com/spf13/cobra
- See https://github.com/hanneshayashi/gsm/tree/main/third_party_licenses for a full list of third party licenses

## See Also

* [Setup](https://gsm.hayashi-ke.online/setup)       - How to set up GSM
* [GSM](https://gsm.hayashi-ke.online/gsm)   - Command overview
* [Batch commands](https://gsm.hayashi-ke.online/batch_commands)     - How to use batch commands
* [Recursive commands](https://gsm.hayashi-ke.online/recursive_commands)     - How to use recursive commands
* [Examples](https://gsm.hayashi-ke.online/examples) - See some examples
* [Scripting examples](https://gsm.hayashi-ke.online/scripting) - Some examples on how to use GSM in scripts
* [PowerShell module](https://github.com/hanneshayashi/gsm-powershell) - Auto-generated PowerShell module
