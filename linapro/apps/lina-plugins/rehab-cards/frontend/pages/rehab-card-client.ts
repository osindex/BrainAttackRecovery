import { requestClient } from '#/api/request';

export interface RehabCategory {
  id: number;
  name: string;
  code: string;
  sort: number;
  status: number;
  remark: string;
}

export interface RehabCard {
  id: number;
  categoryId: number;
  categoryName: string;
  title: string;
  label: string;
  imageUrl: string;
  source: string;
  license: string;
  difficulty: number;
  status: number;
  sort: number;
  remark: string;
}

export interface RehabCrawlerJob {
  id: number;
  categoryId: number;
  keyword: string;
  provider: string;
  requestedCount: number;
  fetchedCount: number;
  status: string;
  message: string;
}

export async function categoryList() {
  const res = await requestClient.get<{ list: RehabCategory[]; total: number }>('/rehab/card/category', { params: { pageNum: 1, pageSize: 100 } });
  return res.list;
}

export function categoryCreate(data: Partial<RehabCategory>) {
  return requestClient.post('/rehab/card/category', data);
}

export function categoryUpdate(id: number, data: Partial<RehabCategory>) {
  return requestClient.put(`/rehab/card/category/${id}`, data);
}

export function categoryDelete(id: number) {
  return requestClient.delete(`/rehab/card/category/${id}`);
}

export async function cardList() {
  const res = await requestClient.get<{ list: RehabCard[]; total: number }>('/rehab/card', { params: { pageNum: 1, pageSize: 100 } });
  return res.list;
}

export function cardCreate(data: Partial<RehabCard>) {
  return requestClient.post('/rehab/card', data);
}

export function cardUpdate(id: number, data: Partial<RehabCard>) {
  return requestClient.put(`/rehab/card/${id}`, data);
}

export function cardDelete(id: number) {
  return requestClient.delete(`/rehab/card/${id}`);
}

export async function crawlerJobList() {
  const res = await requestClient.get<{ list: RehabCrawlerJob[]; total: number }>('/rehab/card/crawler', { params: { pageNum: 1, pageSize: 100 } });
  return res.list;
}

export function crawlerRun(data: { categoryId: number; keyword: string; provider: string; count: number }) {
  return requestClient.post('/rehab/card/crawler', data);
}
