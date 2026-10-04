import { IRepository } from './IRepository';

export class InMemoryRepository implements IRepository {

  private map = new Map<string, string>();

  async get(key: string): Promise<string> {
    return this.map.get(key) ?? null;
  }

  async set(key: string, value: string): Promise<string> {
    this.map.set(key, value);
    return value;
  }

  async remove(key: string): Promise<void> {
    this.map.delete(key);
  }
}