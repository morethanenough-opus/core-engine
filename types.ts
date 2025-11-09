// types.ts
export interface Vector2 {
  x: number;
  y: number;
}

export interface Vector3 {
  x: number;
  y: number;
  z: number;
}

export interface Rectangle {
  x: number;
  y: number;
  width: number;
  height: number;
}

export interface Circle {
  x: number;
  y: number;
  radius: number;
}

export type Shape = Rectangle | Circle;

export interface Collision {
  shape1: Shape;
  shape2: Shape;
}

export enum LogLevel {
  DEBUG,
  INFO,
  WARN,
  ERROR,
}

export interface LogMessage {
  level: LogLevel;
  message: string;
  timestamp: Date;
}