<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { Button, Card, Form, Input, InputNumber, message, Select, Space, Table } from 'ant-design-vue';
import { categoryList, crawlerJobList, crawlerRun, type RehabCategory, type RehabCrawlerJob } from './rehab-card-client';

const loading = ref(false);
const jobs = ref<RehabCrawlerJob[]>([]);
const categories = ref<RehabCategory[]>([]);
const form = reactive({ categoryId: 0, keyword: '', provider: 'manual-placeholder', count: 5 });

async function load() {
  loading.value = true;
  try {
    const [jobRows, categoryRows] = await Promise.all([crawlerJobList(), categoryList()]);
    jobs.value = jobRows;
    categories.value = categoryRows;
    if (!form.categoryId && categoryRows[0]) form.categoryId = categoryRows[0].id;
  } finally { loading.value = false; }
}

async function run() {
  await crawlerRun(form);
  message.success('已生成占位图卡');
  await load();
}

onMounted(load);
</script>

<template>
  <Card title="图卡爬取 / 占位生成">
    <Form layout="inline" style="margin-bottom: 16px">
      <Form.Item label="分类"><Select v-model:value="form.categoryId" style="width: 140px" :options="categories.map((item) => ({ label: item.name, value: item.id }))" /></Form.Item>
      <Form.Item label="关键词"><Input v-model:value="form.keyword" placeholder="苹果" /></Form.Item>
      <Form.Item label="数量"><InputNumber v-model:value="form.count" :min="1" :max="50" /></Form.Item>
      <Form.Item><Button type="primary" :disabled="!form.categoryId || !form.keyword" @click="run">生成图卡</Button></Form.Item>
    </Form>
    <Table :loading="loading" :data-source="jobs" row-key="id" :pagination="false">
      <Table.Column title="关键词" data-index="keyword" />
      <Table.Column title="来源" data-index="provider" />
      <Table.Column title="请求数" data-index="requestedCount" />
      <Table.Column title="完成数" data-index="fetchedCount" />
      <Table.Column title="状态" data-index="status" />
      <Table.Column title="消息" data-index="message" />
    </Table>
  </Card>
</template>
