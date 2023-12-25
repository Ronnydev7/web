interface Logger {
  logError(err: unknown): void;
}

class ConsoleLogger implements Logger {
  constructor(private name: string) {}

  logError(err: unknown): void {
    console.error({
      name: this.name,
      error: err,
    });
  }
}

export function createDefaultLogger(name: string): Logger {
  return new ConsoleLogger(name);
}