import intl from '../intl';

export default function unknownServerError(): string {
  return intl.get('common-strings.unknown-server-error').d('Unknown Server Error');
}