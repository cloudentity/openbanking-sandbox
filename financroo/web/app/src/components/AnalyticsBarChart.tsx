import React, {useState} from 'react';
import {Bar, BarChart, CartesianGrid, ResponsiveContainer, XAxis, YAxis, Cell} from 'recharts';

const data = [
  {
    name: 'JAN', uv: 4000, pv: 2400, amt: 2400,
  },
  {
    name: 'FEB', uv: 3000, pv: 1398, amt: 2210,
  },
  {
    name: 'MAR', uv: 2000, pv: 9800, amt: 2290,
  },
  {
    name: 'APR', uv: 2780, pv: 3908, amt: 2000,
  },
  {
    name: 'MAY', uv: 1890, pv: 4800, amt: 2181,
  },
  {
    name: 'JUN', uv: 2390, pv: 3800, amt: 2500,
  },
  {
    name: 'JUL', uv: 3490, pv: 4300, amt: 2100,
  },
  {
    name: 'AUG', uv: 2000, pv: 9800, amt: 2290,
  },
  {
    name: 'SEP', uv: 2780, pv: 3908, amt: 2000,
  },
  {
    name: 'OCT', uv: 1890, pv: 4800, amt: 2181,
  },
  {
    name: 'NOV', uv: 2390, pv: 3800, amt: 2500,
  },
  {
    name: 'DEC', uv: 3490, pv: 4300, amt: 2100,
  }
];

export default function AnalyticsBarChart() {
  const jsfiddleUrl = 'https://jsfiddle.net/alidingling/q4eonc12/';
  const [selected, setSelected] = useState<any>(null);

  return (
    <ResponsiveContainer width={'100%'} height={300}>
      <BarChart
        data={data}
        margin={{
          top: 5, right: 30, left: 20, bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3"/>
        <XAxis dataKey="name" axisLine={false} tickLine={false}/>
        <YAxis axisLine={false} tickLine={false}/>
        {/*<Tooltip/>*/}
        {/*<Legend/>*/}
        <Bar
          dataKey="pv"
          background={{fill: '#eee'}}
          onClick={e => e.name !== selected?.name
            ? setSelected(e)
            : setSelected(null)}
        >
          {data.map((entry, index) => (
            <Cell cursor="pointer" fill={selected && (entry.name === selected?.name) ? '#36C6AF' : '#1F2D48'} key={`cell-${index}`}/>
          ))}
        </Bar>
      </BarChart>
    </ResponsiveContainer>
  )
}
