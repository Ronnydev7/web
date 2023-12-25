import intl from '../intl';

export default function invalidJwtToken(): string {
  return intl.get('common-strings.invalid-jwt-token').d('Invalid Token');
}