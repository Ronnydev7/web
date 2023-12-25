import isBrowser from '../isBrowser';
import {createDefaultLogger} from '../logger';

import {InMemoryCache} from '@apollo/client';

const LOGGER_NAME = 'lib/createCacheAndRestoreOnce';

export default function createCacheAndRestoreOnce(): InMemoryCache {
  let cache = new InMemoryCache();

  try {
    if (isBrowser && window.INITIAL_QUERY_STATE != null) {
      cache = cache.restore(window.INITIAL_QUERY_STATE);
      window.INITIAL_QUERY_STATE = null;
      const scriptNode = document.getElementById('initial-state');
      if (scriptNode != null) {
        scriptNode.remove();
      }
    }
  } catch (err) {
    createDefaultLogger(LOGGER_NAME).logError(err);
  }
  return cache;
}