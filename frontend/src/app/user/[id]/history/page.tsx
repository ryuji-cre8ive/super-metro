"use client";
import React, { useState, useEffect, useCallback } from "react";
import RecentTransactions from "./components/RecentTransactions";
import axios from "@/api/axiosConfig";
import { useAuth } from "@/app/userContext";
import { Transaction } from "@/app/models/transaction";

const History = () => {
  const { user } = useAuth();

  const [transactions, setTransactions] = useState<Transaction[]>();
  const getTransactions = useCallback(async () => {
    try {
      axios.get(`/transaction/${user?.id}`).then((res) => {
        setTransactions(res.data);
      });
    } catch (e) {
      console.log(e);
    }
  }, [user]);

  useEffect(() => {
    if (!user) return;
    getTransactions();
  }, [user, getTransactions]);

  return (
    <>{transactions && <RecentTransactions transactions={transactions} />}</>
  );
};

export default History;
