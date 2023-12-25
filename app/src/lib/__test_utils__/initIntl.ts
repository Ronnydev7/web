import enUs from '../../generated/translations/en-US.json';

import intl from 'react-intl-universal';

export default function initIntl(): void {
  intl.init({
    currentLocale: 'en-US',
    locales: {'en-US': enUs},
  });
}