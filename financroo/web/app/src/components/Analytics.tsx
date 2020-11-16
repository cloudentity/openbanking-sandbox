import React from "react";
import {Card, Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import AnalyticsTable, {mapTransactionToData} from "./AnalyticsTable";
import AnalyticsBarChart from "./AnalyticsBarChart";
import AnalyticsPieChart from "./AnalyticsPieChart";
import {applyFiltering, mapTransactionsToBarChartData} from "./analytics.utils";
import {pick} from "ramda";

const useStyles = makeStyles((theme: Theme) => ({}));

export default ({transactions, filtering, onChangeFiltering}) => {
  const classes = useStyles();

  const barChartData = mapTransactionsToBarChartData(applyFiltering(pick(['accounts'], filtering), transactions));
  const pieChartData = applyFiltering(pick(['accounts', 'months'], filtering), transactions);
  const tableData = applyFiltering(pick(['accounts', 'months', 'categories'], filtering), transactions).map(mapTransactionToData);

  return (
    <>
      <Card style={{padding: 16, display: 'flex', alignItems: 'center'}}>
        <div style={{flex: 3}}>
          <AnalyticsBarChart data={barChartData} filtering={filtering} onChangeFiltering={onChangeFiltering}/>
        </div>
        <div style={{flex: 1}}>
          <AnalyticsPieChart data={pieChartData} filtering={filtering} onChangeFiltering={onChangeFiltering}/>
        </div>
      </Card>
      <AnalyticsTable data={tableData} style={{marginTop: 24, height: 'calc(100% - 332px - 24px'}}/>
    </>
  )
};
