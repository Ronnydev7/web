import App from './App';
import {AppContextValue} from './lib/AppContext';

import {loadableReady} from '@loadable/component';
import React from 'react';
import ReactDOM from 'react-dom/client';

async function init(): Promise<void> {
  const initialContext: AppContextValue = {};

  ReactDOM.hydrateRoot(
    document.getElementById('root') as Element,
    <App initialContext={initialContext} />,
  );
}

loadableReady(init);