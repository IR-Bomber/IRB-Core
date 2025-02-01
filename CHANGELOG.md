## 1.0.0 (2025-02-01)

### 🚀 Features

* Add support for multiple architectures in CI build workflow ([a2b0fc4](https://github.com/IR-Bomber/IRB-Core/commit/a2b0fc4354db4b3b46cc295e7a328f91b2e30379))
* Enhance API fetching and validation with improved error handling and testing ([a060e66](https://github.com/IR-Bomber/IRB-Core/commit/a060e66c5da6baced3dedc4ff541bf9a67ed1ba9))
* Enhance release workflow with semantic-release integration and CalVer plugin ([35e0bf7](https://github.com/IR-Bomber/IRB-Core/commit/35e0bf7b7036108000417db6d4cece1920ffccff))

### 🐛 Bug Fixes

* Add CalVer plugin to release configuration ([df6d83e](https://github.com/IR-Bomber/IRB-Core/commit/df6d83e92f329527b0fa35ecda70cccded61110d))
* Adjust build script for Windows compatibility in CI workflow ([1cebe09](https://github.com/IR-Bomber/IRB-Core/commit/1cebe091c225f46ca73c6d138a1882b570df2a1e))
* Correct build command syntax in CI workflow ([766b53d](https://github.com/IR-Bomber/IRB-Core/commit/766b53d41745541ba45e9e58e8d9fd99e9f88561))
* Correct plugin path for CalVer in release configuration ([e46ef7d](https://github.com/IR-Bomber/IRB-Core/commit/e46ef7d1296aba1644942b3926a10a298baed2ae))
* Disable fail-fast in CI workflow to allow all jobs to complete ([3cd1e68](https://github.com/IR-Bomber/IRB-Core/commit/3cd1e68d2840bb8aea98c6b9078f14b9c4d7549c))
* Enhance CI workflow to support Go tests and builds on Windows ([4b6dde1](https://github.com/IR-Bomber/IRB-Core/commit/4b6dde1d9bd82c792c214cb25990ff70b8ee8105))
* Ensure all binaries are moved to dist directory in CI workflow ([3d1145e](https://github.com/IR-Bomber/IRB-Core/commit/3d1145e0d09702e5087c67c7d9564cecc9ff8e60))
* Ensure build job depends on test job in CI workflow ([6857f54](https://github.com/IR-Bomber/IRB-Core/commit/6857f5462359d814c13696769a18422783b2dab9))
* Remove duplicate comment in TestAPI function ([76549a2](https://github.com/IR-Bomber/IRB-Core/commit/76549a2c2bb960bfffc3439f6396430b1f78c42c))
* Separate Go test execution for Linux and MacOS in CI workflow ([161abdf](https://github.com/IR-Bomber/IRB-Core/commit/161abdfae48de84ff0dcaa4cb608bd3616f7d8f3))
* Simplify Go build and test commands in CI workflow ([7ddbf6d](https://github.com/IR-Bomber/IRB-Core/commit/7ddbf6d2c3bf43856edcb2b49f8056af199ea3db))
* Update API URL to point to the correct data path ([2dc6b66](https://github.com/IR-Bomber/IRB-Core/commit/2dc6b66fff8b9b5fd272455a516a6541c68e7930))
* Update API URL to point to the correct repository ([6e8ec28](https://github.com/IR-Bomber/IRB-Core/commit/6e8ec286c2286949d2b09ab641dfd012b7522713))
* Update binary artifact naming to include OS in CI workflow ([b20aec8](https://github.com/IR-Bomber/IRB-Core/commit/b20aec85f118f84151eb29ba25aff6e261ba35ab))
* Update CI workflow to run Go tests inside Docker for Linux/MacOS ([0a3024d](https://github.com/IR-Bomber/IRB-Core/commit/0a3024d53457ed46bb1d5708e0f780bba5261e7d))
* Update CI workflow to set up Docker only for Linux runners ([f40a817](https://github.com/IR-Bomber/IRB-Core/commit/f40a8178159103bc55cdbf838ef2149be3379950))
* Update condition for dry-run checks in release workflow ([fe18163](https://github.com/IR-Bomber/IRB-Core/commit/fe18163474d112cdce6142612648df6f3a1575f1))
* Update import path to reflect new repository name ([4dc430e](https://github.com/IR-Bomber/IRB-Core/commit/4dc430e2b4557f112aa7155b5619d26627f13211))
* Update module path and import statements to reflect new repository name ([a0f84cb](https://github.com/IR-Bomber/IRB-Core/commit/a0f84cb92a62c86df014c3fcd556a533b12b97e9))
* Update release branch from master to main ([aa25c08](https://github.com/IR-Bomber/IRB-Core/commit/aa25c08b1115baf0fd3698e5518019fb941ce2b4))
* Update release workflow permissions for better access control ([4c4126d](https://github.com/IR-Bomber/IRB-Core/commit/4c4126d75551f3542ad16d99859c542bc7a4b012))

### 🚧 Refactors

* Modularize CI workflow by separating test execution into its own file ([4e2584b](https://github.com/IR-Bomber/IRB-Core/commit/4e2584bf25086559573b2c254d2d499e6e1fd5d2))
* Move API fetching logic to utils package and update references ([062a0f8](https://github.com/IR-Bomber/IRB-Core/commit/062a0f8b360ac19320955ef15ed1dd76fc712084))

### 🔄 CI/CD

* Add CI and build workflows for automated testing and binary generation ([dc33f84](https://github.com/IR-Bomber/IRB-Core/commit/dc33f845fb530931232c802eb088d739cd870503))
