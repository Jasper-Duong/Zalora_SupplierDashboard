import { useState, useEffect } from "react";
import { Table, Tag } from "antd";
import qs from "qs";
import { request } from "../../config/axios";

const ExtendedTable = (props) => {
  const [data, setData] = useState(props.data);
  const [tableParams, setTableParams] = useState({
    pagination: {
      showQuickJumper: true,
      pageSizeOptions: ["10", "50", "100"],
      showSizeChanger: true,
      current: 1,
      pageSize: 10,
    },
    sorter: [],
  });

  const getQueryParams = (params) => {
    let filters = {};
    if (params.filters) {
      filters = Object.entries(params.filters).reduce((prev, [key, value]) => {
        if (value !== null) prev[key] = value;
        return prev;
      }, {});
    }
    return {
      page: params.pagination?.current,
      limit: params.pagination?.pageSize,
      sort: params.sorter.map((sort) => sort.columnKey),
      order: params.sorter.map((sort) =>
        sort.order === "ascend" ? "asc" : "desc"
      ),
      ...filters,
    };
  };

  const handleTableChange = (pagination, filters, sorter) => {
    setTableParams({
      pagination,
      filters,
      sorter: Array.isArray(sorter) ? sorter : [sorter],
    });
  };

  useEffect(() => {
    async function fetchData() {
      try {
        const res = await request.get(
          `${props.resource}/?${decodeURIComponent(
            qs.stringify(getQueryParams(tableParams), {
              arrayFormat: "brackets",
            })
          )}`
        );
        setData(res.data[`${props.resource}`]);
        setTableParams({
          ...tableParams,
          pagination: {
            ...tableParams.pagination,
            total: res.data.total,
          },
        });
      } catch (err) {
        console.log(err);
      }
    }
    fetchData();
  }, [JSON.stringify(tableParams)]);

  return (
    <Table
      dataSource={data}
      columns={props.columns}
      pagination={tableParams.pagination}
      onChange={handleTableChange}
      rowKey="id"
      expandable={{
        expandedRowRender: (record) => (
          <Tag
            style={{
              margin: "auto",
            }}
          >
            {"cuu"}
          </Tag>
        ),
      }}
    ></Table>
  );
};

export default ExtendedTable;
