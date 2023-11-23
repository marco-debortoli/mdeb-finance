import type { Account } from "@/types/account";
import dayjs from "dayjs";

export function getActiveAccountsForDate(accounts: Account[], date: Date) {
  const currentDate = dayjs(date);

  return accounts.filter((acc) => {
    const accStartDate = dayjs(acc.start_date).startOf('month');

    if (accStartDate.isAfter(currentDate)) return false;
    if (!acc.end_date) return true;

    const accEndDate = dayjs(acc.end_date).startOf('month');
    return accEndDate.isSame(currentDate) || accEndDate.isAfter(currentDate);
  });
};

export function getAccountsTotalForDate(accounts: Account[], date: Date) {
  let total = 0;
  const currentDate = dayjs(date);

  accounts.forEach((a) => {
    const av = a.values.find((v) => currentDate.isSame(dayjs(v.date)));

    if (av) total += av.value;
  });

  return total;
};

export function absoluteAccountValueDiff(v1: number | undefined, v2: number | undefined) {
  let diff = 0;

  if (v1 !== undefined && v2 !== undefined) {
    diff = v1 - v2;
  } else if (v1 !== undefined && v2 === undefined) {
    diff = v1;
  } else if (v1 === undefined && v2 !== undefined) {
    diff = -v2;
  } else {
    diff = 0
  }

  return diff
};