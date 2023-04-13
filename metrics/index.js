import { CloudWatchClient, PutMetricDataCommand } from "@aws-sdk/client-cloudwatch";
import { CloudWatchLogsClient, PutLogEventsCommand, CreateLogStreamCommand } from "@aws-sdk/client-cloudwatch-logs";
import os from "node:os";

const cloudwatchClient = new CloudWatchClient({region: "ap-southeast-2"});
const clientwatchLogClient = new CloudWatchLogsClient({region: "ap-southeast-2"});

async function main() {
  // await clientwatchLogClient.send(new CreateLogStreamCommand({ // CreateLogStreamRequest
  //   logGroupName: "/CWAgent/Fergus/Queue/syslog", // required
  //   logStreamName: os.hostname(), // required
  // }));

  await seedErrorLogs();

  const dataArray = [
    {
      supplierName: "Reece",
      invoiceSource: "drag_and_drop",
      value: 30,
      timestamp: new Date("2023-04-12T12:34:00Z")
    },
    {
      supplierName: "Reece",
      invoiceSource: "email",
      value: 5,
      timestamp: new Date("2023-04-13T01:23:00Z")
    },
    {
      supplierName: "Reece",
      invoiceSource: "premium_or_unknown_source",
      value: 50,
      timestamp: new Date("2023-04-13T01:23:00Z")
    }
    ];

    putMetric(dataArray);
}

main()

async function seedErrorLogs() {
  const rand = (min, max) => {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1) + min); // The maximum is inclusive and the minimum is inclusive
  }

  const supplierNames = ["Reece", "Mico", "JA Russel", "Tradelink"];

  const input = {
    // PutLogEventsRequest
    logGroupName: "/CWAgent/Fergus/Queue/syslog", // required
    logStreamName: os.hostname(), // required
    logEvents: [],
  };

  for (let i = 0; i <= 100; i++) {
    const id = rand(0, supplierNames.length-1);
    const batch = rand(0, 100);
    const event = {
      timestamp: Date.now(), // required
      message: JSON.stringify({
        message: "INVOICE_IMPORT_FAILURE", // required
        context: {
          supplier_id: id,
          supplier_name: supplierNames[id],
          company_id: id,
          file: {
            file_id: 1,
            file_size: 1,
            s3_path: "/tmp",
          },
          exception: "Some exception",
          batch_id: batch,
        },
      }),
    }

    input.logEvents.push(event);
  }

  const command = new PutLogEventsCommand(input);

  await clientwatchLogClient.send(command);
}

async function putMetric(dataArray) {

  const params = {
    Namespace: "invoice-ingestion-stats-poc",
    MetricData: dataArray.map((data) => ({
      MetricName: "Total invoices ingested",
      Dimensions: [
        {
          Name: "InvoiceSource",
          Value: data.invoiceSource
        },
        {
          Name: "Supplier",
          Value:  data.supplierName
        }
      ],
      Timestamp: data.timestamp,
      Value: data.value,
      Unit: "Count"
    }))
  };
  await cloudwatchClient.send(new PutMetricDataCommand(params));
}
