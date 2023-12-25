import reactIntlUniversal from 'react-intl-universal';

// eslint-disable-next-line @typescript-eslint/no-unused-vars
declare interface String {
  d(msg: string | JSX.Element): string;
}

export type Intl = Readonly<{
  get(key: string, variables?: Record<string, unknown>): string;
}>;

// This is only intended for react-intl-universal-extract to extract the strings.
const intl: Intl = reactIntlUniversal;

export default intl;