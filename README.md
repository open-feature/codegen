<!-- markdownlint-disable MD033 -->
<!-- x-hide-in-docs-start -->
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/open-feature/community/0e23508c163a6a1ac8c0ced3e4bd78faafe627c7/assets/logo/horizontal/white/openfeature-horizontal-white.svg" />
    <img align="center" alt="OpenFeature Logo" src="https://raw.githubusercontent.com/open-feature/community/0e23508c163a6a1ac8c0ced3e4bd78faafe627c7/assets/logo/horizontal/black/openfeature-horizontal-black.svg" />
  </picture>
</p>

<h2 align="center">OpenFeature CodeGen</h2>
<!-- x-hide-in-docs-end -->
<!-- The 'github-badges' class is used in the docs -->
<p align="center" class="github-badges">
  <a href="https://github.com/orgs/open-feature/projects/17">
    <img alt="work-in-progress" src="https://img.shields.io/badge/status-WIP-red" />
  </a>
  <a href="https://cloud-native.slack.com/archives/C07DY4TUDK6">
    <img alt="Slack" src="https://img.shields.io/badge/slack-%40cncf%2Fopenfeature-brightgreen?style=flat&logo=slack" />
  </a>
</p>
<!-- x-hide-in-docs-start -->

> [!CAUTION]
> OpenFeature CodeGen experimental!
> Feel free to give it a shot and provide feedback but expect breaking changes.

[OpenFeature](https://openfeature.dev) is an open specification that provides a vendor-agnostic, community-driven API for feature flagging that works with your favorite feature flag management tool or in-house solution.
<!-- x-hide-in-docs-end -->

## Why CodeGen?

CodeGen automates the creation of strongly typed flag accessors, minimizing configuration errors and provides a better developer experience.
By generating these accessors, developers can avoid issues related to incorrect flag names or types, resulting in more reliable and maintainable code.

## Goals

- **Unified Flag Manifest Format**: Establish a standardized flag manifest format that can be easily converted from existing configurations.
- **Strongly Typed Flag Accessors**: Develop a CLI tool to generate strongly typed flag accessors for multiple programming languages.
- **Modular and Extensible Design**: Create a format that allows for future extensions and modularization of flags.

## Non-Goals

- **Full Provider Integration**: The initial scope does not include creating tools to convert provider-specific configurations to the new flag manifest format.
- **Validation of Flag Configs**: The project will not initially focus on validating flag configurations for consistency with the flag manifest.
- **General-Purpose Configuration**: The project will not aim to create a general-purpose configuration tool for feature flags beyond the scope of the codegen tool.

## Support the project

- Give this repo a ⭐️!
- Follow us on social media:
  - Twitter: [@openfeature](https://twitter.com/openfeature)
  - LinkedIn: [OpenFeature](https://www.linkedin.com/company/openfeature/)
- Join us on [Slack](https://cloud-native.slack.com/archives/C0344AANLA1)
- For more, check out our [community page](https://openfeature.dev/community/)

### Thanks to everyone who has already contributed

<a href="https://github.com/open-feature/codegen/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=open-feature/codegen" alt="Pictures of the folks who have contributed to the project" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
<!-- x-hide-in-docs-end -->
