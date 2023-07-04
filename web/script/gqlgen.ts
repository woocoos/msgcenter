import { CodegenConfig } from "@graphql-codegen/cli";
import * as process from "process";

const dotenv = require('dotenv')
dotenv.config()
dotenv.config({ path: `.env.${process.env.NODE_ENV}`, override: true })
dotenv.config({ path: '.env.local', override: true })

const adminxSchema = process.env.GQLGEN_SCHEMA_ADMINX

if (!adminxSchema) {
  throw Error('The env.GQLGEN_SCHEMA_ADMINX is undefined')
}

/**
 * 生成.graphql的配置
 */
const schemaAstConfig: CodegenConfig = {
  generates: {
    // adminx 项目
    'script/__generated__/adminx.graphql': {
      plugins: ['schema-ast'],
      config: {
        includeDirectives: true,
      },
      schema: {
        [adminxSchema]: {
          headers: {
            "Authorization": `Bearer ${process.env.GQLGEN_TOKEN}`,
            "X-Tenant-ID": `${process.env.GQLGEN_TENANT_ID}`,
          }
        },
      }
    },
    // msgsrv
    'script/__generated__/msgsrv.graphql': {
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
    // adminx 项目
    "src/__generated__/adminx/": {
      preset: 'client',
      presetConfig: {
        gqlTagName: 'gql',
      },
      schema: "script/__generated__/adminx.graphql",
      documents: "src/services/adminx/**/*.ts",
    },
    // msgsrv 项目
    "src/__generated__/msgsrv/": {
      preset: 'client',
      presetConfig: {
        gqlTagName: 'gql',
      },
      schema: "script/__generated__/msgsrv.graphql",
      documents: "src/services/msgsrv/**/*.ts",
    }
  },
  ignoreNoDocuments: true,
}


export default process.argv.includes('--schema-ast') ? schemaAstConfig : config
