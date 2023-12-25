
import type {CodegenConfig} from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: '../graphql/schema.json',
  documents: ['src/**/*.ts*', '!src/**/*.graphql.ts*'],
  generates: {
    'src/graphql/types.ts': {
      plugins: ['typescript'],
    },
    'src/': {
      preset: 'near-operation-file',
      presetConfig: {
        extension: '.graphql.ts',
        baseTypesPath: 'graphql/types.ts',
      },
      plugins: [
        'typescript-operations', 
        'typescript-react-apollo',
      ],
    },
  },
};

export default config;
