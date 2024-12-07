# kcommit
Git commit generator on Karma style


## kcommit custom configs
By default kcommit used the basic types from karma such as:

- feat: Adding a new feature.
- fix: Fixing a bug.
- refactor: Code change that does not change functionality or fix bugs.
- docs: Documentation update.
- style: Changes related to code style (spaces, missing semicolons, etc.) without affecting the logic.
- test: Adding or updating tests.
- chore: Updates that do not affect production code (build configuration, dependencies, etc.).

But it's possible to define your own types creating a file `.kcommitrc` at the root of your project following the format below:

```json
{
  "commitTypes": [
    {
      "type": "feat",
      "description": "Adds a new feature to the project.",
      "example": "feat(auth): add support for Google authentication"
    },
    {
      "type": "fix",
      "description": "Fixes a bug in the code.",
      "example": "fix(payment): fix discount calculation error"
    },
    {
      "type": "docs",
      "description": "Updates documentation only, without changing the code.",
      "example": "docs(readme): add instructions to initialize the project"
    },
    {
      "type": "style",
      "description": "Changes that do not affect functionality (e.g., formatting, whitespace).",
      "example": "style(ui): adjust spacing between buttons in the form"
    },
    {
      "type": "refactor",
      "description": "Refactors code without changing existing functionality.",
      "example": "refactor(api): optimize authentication logic"
    },
    {
      "type": "test",
      "description": "Adds or updates automated tests.",
      "example": "test(api): add tests for the login route"
    },
    {
      "type": "chore",
      "description": "Auxiliary tasks, such as dependency updates or configuration changes.",
      "example": "chore(deps): update ESLint to version 8.4.0"
    },
    {
      "type": "perf",
      "description": "Performance improvements in the code.",
      "example": "perf(render): reduce loading time for the homepage"
    },
    {
      "type": "ci",
      "description": "Changes to the continuous integration configuration.",
      "example": "ci(actions): fix deploy workflow in GitHub Actions"
    },
    {
      "type": "build",
      "description": "Changes related to the build system or external dependencies.",
      "example": "build: add support for builds on ARM environments"
    },
    {
      "type": "revert",
      "description": "Reverts a previous commit.",
      "example": "revert: revert commit abc1234 due to production issues"
    }
  ]
}
```