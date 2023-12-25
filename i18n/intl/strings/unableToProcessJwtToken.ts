import intl from '../intl';

export default function unableToProcessJwtToken(): string {
  return intl.get('common-strings.unable-to-process-jwt-token').d('Unable to process token');
}