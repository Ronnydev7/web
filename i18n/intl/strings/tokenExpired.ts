import intl from '../intl';

export default function expiredToken(): string {
  return intl.get('common-strings.token-expired').d('Token Expired');
}