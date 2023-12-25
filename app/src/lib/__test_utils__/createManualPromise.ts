export type ManualPromiseCallbacks<T> = {
  resolve: (result: T) => void,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  reject: (data: any) => void,
}

export type ManualPromise<T> = {
  promise: Promise<T>,
  callbacks: ManualPromiseCallbacks<T>,
};
export default function createManualPromise<T>(): ManualPromise<T> {
  const callbacks: ManualPromiseCallbacks<T> = {
    resolve: () => undefined,
    reject: () => undefined,
  };
  const promise = new Promise<T>(
    (res, rej) => {
      callbacks.resolve = res;
      callbacks.reject = rej;
    },
  );
  return {
    promise,
    callbacks,
  };
}