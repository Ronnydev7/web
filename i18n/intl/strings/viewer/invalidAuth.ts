import intl from '../../intl';

export default function invalidLoginCredential(): string {
  return intl.get('viewer.invalid-login-credential').d('Invalid Auth Information');
}