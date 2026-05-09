<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { Button, Card, Form, Image, Input, InputNumber, message, Popconfirm, Select, Space, Table } from 'ant-design-vue';
import { cardCreate, cardDelete, cardList, cardUpdate, categoryList, type RehabCard, type RehabCategory } from './rehab-card-client';

const loading = ref(false);
const cards = ref<RehabCard[]>([]);
const categories = ref<RehabCategory[]>([]);
const editingId = ref<number | null>(null);
const form = reactive({ categoryId: 0, title: '', label: '', imageUrl: '', source: 'manual', license: '', difficulty: 1, status: 1, sort: 0, remark: '' });

async function load() {
  loading.value = true;
  try {
    const [cardRows, categoryRows] = await Promise.all([cardList(), categoryList()]);
    cards.value = cardRows;
    categories.value = categoryRows;
    if (!form.categoryId && categoryRows[0]) form.categoryId = categoryRows[0].id;
  } finally { loading.value = false; }
}

function reset() {
  editingId.value = null;
  Object.assign(form, { categoryId: categories.value[0]?.id ?? 0, title: '', label: '', imageUrl: '', source: 'manual', license: '', difficulty: 1, status: 1, sort: 0, remark: '' });
}

function edit(row: RehabCard) {
  editingId.value = row.id;
  Object.assign(form, { categoryId: row.categoryId, title: row.title, label: row.label, imageUrl: row.imageUrl, source: row.source, license: row.license, difficulty: row.difficulty, status: row.status, sort: row.sort, remark: row.remark });
}

async function save() {
  if (editingId.value) await cardUpdate(editingId.value, form);
  else await cardCreate(form);
  message.success('已保存');
  reset();
  await load();
}

async function remove(id: number) {
  await cardDelete(id);
  message.success('已删除');
  await load();
}

onMounted(load);
</script>

<template>
  <Card title="图片卡片管理">
    <Form layout="vertical" style="margin-bottom: 16px">
      <Space wrap>
        <Form.Item label="分类"><Select v-model:value="form.categoryId" style="width: 140px" :options="categories.map((item) => ({ label: item.name, value: item.id }))" /></Form.Item>
        <Form.Item label="标题"><Input v-model:value="form.title" placeholder="苹果" /></Form.Item>
        <Form.Item label="答案"><Input v-model:value="form.label" placeholder="苹果" /></Form.Item>
        <Form.Item label="图片 URL"><Input v-model:value="form.imageUrl" style="width: 320px" placeholder="https://..." /></Form.Item>
        <Form.Item label="难度"><InputNumber v-model:value="form.difficulty" :min="1" :max="3" /></Form.Item>
        <Form.Item label="排序"><InputNumber v-model:value="form.sort" /></Form.Item>
        <Form.Item><Button type="primary" @click="save">{{ editingId ? '更新' : '新增' }}</Button></Form.Item>
        <Form.Item><Button @click="reset">清空</Button></Form.Item>
      </Space>
    </Form>
    <Table :loading="loading" :data-source="cards" row-key="id" :pagination="false">
      <Table.Column title="图片"><template #default="{ record }"><Image :src="record.imageUrl" :width="72" /></template></Table.Column>
      <Table.Column title="标题" data-index="title" />
      <Table.Column title="答案" data-index="label" />
      <Table.Column title="分类" data-index="categoryName" />
      <Table.Column title="难度" data-index="difficulty" />
      <Table.Column title="操作"><template #default="{ record }"><Space><Button size="small" @click="edit(record)">编辑</Button><Popconfirm title="确认删除？" @confirm="remove(record.id)"><Button size="small" danger>删除</Button></Popconfirm></Space></template></Table.Column>
    </Table>
  </Card>
</template>
