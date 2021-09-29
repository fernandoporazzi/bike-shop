import { useEffect, useState } from "react";
import { Bike } from "../../types";
import List from "../../components/List";
import Empty from "../../components/Empty";

type Response = {
  bikes: Bike[];
};

const Home = (): JSX.Element => {
  const [bikes, setBikes] = useState<Bike[]>([]);

  useEffect(() => {
    const request = async () => {
      const res = await fetch("http://localhost:8080/bikes");
      const data: Response = await res.json();

      if (data.bikes) {
        setBikes(data.bikes);
      }
    };

    request();
  }, [])

  return (
    bikes.length > 0 ? <List bikes={bikes} /> : <Empty />
  )
};

export default Home;
