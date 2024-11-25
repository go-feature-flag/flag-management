abstract class GoFeatureFlagError extends Error {
  constructor(message: string) {
    super(message);
    this.name = this.constructor.name;
    Error.captureStackTrace(this, this.constructor);
  }
}

export class ConvertToDTOError extends GoFeatureFlagError {}
