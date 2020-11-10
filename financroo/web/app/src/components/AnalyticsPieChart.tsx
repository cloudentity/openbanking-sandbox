import React from 'react';
import {Cell, Pie, PieChart, ResponsiveContainer} from 'recharts';
import {Typography} from "@material-ui/core";
import classes from "./chartsStyles.module.css";

const data = [
  {name: 'Group A', value: 400},
  {name: 'Group B', value: 300},
  {name: 'Group C', value: 300},
  {name: 'Group D', value: 200},
];
const COLORS = ['#0088FE', '#00C49F', '#FFBB28', '#FF8042'];

export default function AnalyticsPieChart() {

  return (
    <div style={{position: 'relative'}}>
      <div className={classes.pieChartContent}>
        <Typography style={{fontSize: 14, fontWeight: 600, color: '#626576'}}>October</Typography>
        <Typography style={{fontSize: 16, fontWeight: 600, marginTop: 6}}>€ 4,500.00</Typography>
        <Typography style={{fontSize: 12, marginTop: 2, color: '#626576'}}>Entertainment</Typography>
      </div>
      <ResponsiveContainer width={'100%'} height={300}>
        <PieChart>
          <Pie
            data={data}
            // cx={120}
            // cy={200}
            innerRadius={90}
            outerRadius={120}
            fill="#8884d8"
            // paddingAngle={5}
            dataKey="value"
          >
            {
              data.map((entry, index) => <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]}/>)
            }
          </Pie>
        </PieChart>
      </ResponsiveContainer>
    </div>
  )
}
