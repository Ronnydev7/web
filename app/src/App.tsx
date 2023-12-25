import type {AppContextValue} from './lib/AppContext';

import AppContext from './lib/AppContext';

import React from 'react';

type Props = Readonly<{
  initialContext: AppContextValue,
}>;

export default function App(props: Props): JSX.Element {
  const {initialContext} = props;

  return (
    <AppContext.Provider value={initialContext}>
      <div>TODO</div>
    </AppContext.Provider>
  );
}