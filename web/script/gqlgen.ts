import { CodegenConfig } from "@graphql-codegen/cli";
import * as process from "process";

const dotenv = require('dotenv')
dotenv.config()
dotenv.config({ path: '.env.local', override: true })

/**
 * 生成.graphql的配置
 */
const schemaAstConfig: CodegenConfig = {
  generates: {
    // msgsrv
    'script/generated/msgsrv.graphql': {
      plugins: ['schema-ast'],
      config: {
        includeDirectives: true,
      },
      schema: "../api/graphql/*.graphql"
    }
  }
}


/**
 * 开发使用的生成配置
 */
const config: CodegenConfig = {
  generates: {
    // msgsrv 项目
    "src/generated/msgsrv/": {
      preset: 'client',
      presetConfig: {
        gqlTagName: 'gql',
      },
      schema: "script/generated/msgsrv.graphql",
      documents: "src/services/msgsrv/**/*.ts",
    }
  },
  ignoreNoDocuments: true,
}


export default process.argv.includes('--schema-ast') ? schemaAstConfig : config
