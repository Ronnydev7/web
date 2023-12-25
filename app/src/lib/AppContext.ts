import React from 'react';

export type AppContextValue = Readonly<{
  statusCode?: number,
}>;

export default React.createContext<AppContextValue>({});