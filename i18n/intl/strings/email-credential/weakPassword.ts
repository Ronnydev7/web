import intl from '../../intl';

export default function weakPassword(): string {
  return intl.get('email-credential.weak-password').d('Password is too weak. We use entropy to detect weak password. Please add more variance to your password characters');
}