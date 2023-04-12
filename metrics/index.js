import { CloudWatchClient, PutMetricDataCommand } from "@aws-sdk/client-cloudwatch";

const cloudwatchClient = new CloudWatchClient({region: "ap-southeast-2"});

async function putMetricReece() {
  const params = {
  Namespace: "invoice-ingestion-stats-poc",
  MetricData: [
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "manual"
        },
        {
          Name: "Supplier",
          Value: "Reece"
        }
      ],
      Timestamp: new Date(),
      Value: 10,
      Unit: "Count"
    },
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "drag_and_drop"
        },
        {
          Name: "Supplier",
          Value: "Reece"
        }
      ],
      Timestamp: new Date(),
      Value: 30,
      Unit: "Count"
    },
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "email"
        },
        {
          Name: "Supplier",
          Value: "Reece"
        }
      ],
      Timestamp: new Date(),
      Value: 5,
      Unit: "Count"
    },
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "premium_or_unknown_source"
        },
        {
          Name: "Supplier",
          Value: "Reece"
        }
      ],
      Timestamp: new Date(),
      Value: 50,
      Unit: "Count"
    }
  ]
};
  try {

    const data = await cloudwatchClient.send(new PutMetricDataCommand(params));
    return data;
  } catch (err) {
    console.log("Error sending metric:", err);
  }
}

async function putMetricMico() {
  const params = {
  Namespace: "invoice-ingestion-stats-poc",
  MetricData: [
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "manual"
        },
        {
          Name: "Supplier",
          Value: "Mico"
        }
      ],
      Timestamp: new Date(),
      Value: 11,
      Unit: "Count"
    },
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "drag_and_drop"
        },
        {
          Name: "Supplier",
          Value: "Mico"
        }
      ],
      Timestamp: new Date(),
      Value: 33,
      Unit: "Count"
    },
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "email"
        },
        {
          Name: "Supplier",
          Value: "Mico"
        }
      ],
      Timestamp: new Date(),
      Value: 4,
      Unit: "Count"
    },
    {
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: "premium_or_unknown_source"
        },
        {
          Name: "Supplier",
          Value: "Mico"
        }
      ],
      Timestamp: new Date(),
      Value: 1,
      Unit: "Count"
    }
  ]
};
  try {

    const data = await cloudwatchClient.send(new PutMetricDataCommand(params));
    return data;
  } catch (err) {
    console.log("Error sending metric:", err);
  }
}

putMetricReece();

putMetricMico();
