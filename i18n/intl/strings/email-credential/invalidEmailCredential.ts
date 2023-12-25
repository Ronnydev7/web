import intl from '../../intl';

export default function invalidLoginCredential(): string {
  return intl.get('email-credential.invalid-email-credential')
    .d('Invalid Login Credential');
}