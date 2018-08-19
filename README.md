## Au Fait

The NPM dependency management tool for everyone! This is still under development. Please watch our progress. Projected release date coming soon!

### Development (I live a lonely life)

Any feedback, ideas, requests can be sent to the public email address associated with this project.

### How it works (I hope)

Au Fait is a collection of two micro services. A simple listener and a build. We listen for Git and NPM updates and we update based on repository configuration. Each repository is updated, tests are run, and we notify you of the outcome. In both success and failure scenarios we push to a branch and can be configured to open a pull request.

### Goals (hashtag gooooals)

We are building a GreenKeeper.io (cap tipped) like tool that works with more services. We want to support self hosted and publicly hosted setups. We also support the below services (and are always opened to pluging in more services).

### Supported Version Control Services

* GitHub
* GitLab
* Stash/BitBucket

### Supported Messaging Services

* HipChat
* Slack

### Supported NPM registries

* Public NPM
* Configured Artifactory NPM registries

### Getting Started (under development)

* [Install the aufait command line utility](https://github.com/quantumew/aufait/releases).
* Point it to repositories it should manage. It will setup Git WebHooks to ensure it stays up to date!
* Follow prompts to configure repository (this step is skipped when config is present).
* Setup any private NPM registries. (Instructions to come)
