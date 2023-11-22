import type { Category } from "@/types/category";

export type Transaction = {
    _id: string;
    amount: number;
    date: string;
    name: string;
    category: Category;
}