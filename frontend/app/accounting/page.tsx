"use client";

import React, {useState} from 'react';
import withAuth from '../hocs/withAuth';
import Accounts from '../components/Accounts';
import Transaction from '../components/Transaction';
import MainLayout from '../components/MainLayout';
import ContentHeading from '../components/ContentHeading';


enum AccountingKeys {
    accounts = "accounts",
    transaction = "transaction",
  }
type pageProps = {
    
};

const Accounting = () => {
    const [activeTab, setActiveTab] = useState(AccountingKeys.accounts);
    const accountingComponents = {
        [AccountingKeys.accounts]: <Accounts />,
        [AccountingKeys.transaction]: <Transaction />,
      };
    
    return (
        <MainLayout>
            <main>
            <ContentHeading
                title="Accounting"
                sideSection={
                     <SideSection activeTab={activeTab} onSetActiveTab={setActiveTab} />
                }
            />
            {accountingComponents[activeTab]}
            </main>
        </MainLayout>
    )
}

interface SideSectionType {
    activeTab: string;
    onSetActiveTab: (tab: AccountingKeys) => void;
  }


  const SideSection = (props: SideSectionType) => {
    const getIsActive = (tab: string) => {
      if (props.activeTab === tab) return "active";
      return "";
    };
  
    return (
      <div className="sideTab">
        <div
          className={`item ${getIsActive("accounts")}`}
          onClick={() => props.onSetActiveTab(AccountingKeys.accounts)}
        >
          Accounts
        </div>
        <div
          className={`item ${getIsActive("transaction")}`}
          onClick={() => props.onSetActiveTab(AccountingKeys.transaction)}
        >
          Transaction
        </div>
      </div>
    );
  };
export default withAuth(Accounting);