import { Layout } from "antd";
import { useState } from "react";
import { Outlet } from "react-router-dom";
import HomeMenu from "../../components/HomeMenu/HomeMenu";
import HomeSiderLogo from "./HomeSiderLogo";
const { Content, Sider } = Layout;

const HomeLayout = () => {
  const [collapsed, setCollapsed] = useState(false);
  return (
    <Layout
      style={{
        minHeight: "100vh",
      }}
    >
      <Sider
        collapsible
        collapsed={collapsed}
        onCollapse={(value) => setCollapsed(value)}
      >
        <HomeSiderLogo />
        <HomeMenu />
      </Sider>
      <Layout className="site-layout">
        <Content
          style={{
            margin: "0 16px",
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};
export default HomeLayout;
