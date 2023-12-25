import intl from 'react-intl-universal';

export default function authorized(): string {
  return intl.get('common-strings.unauthorized')
    .d('Unauthorized Access');
}