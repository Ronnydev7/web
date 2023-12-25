import intl from '../../intl';

export default function duplicateHandleName(): string {
  return intl.get('public-profile.duplicate-handle-name').d('This handle name has been taken');
}