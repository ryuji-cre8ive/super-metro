/**
 * This code was generated by v0 by Vercel.
 * @see https://v0.dev/t/WyjRNvCq7tP
 */
import {
  CardTitle,
  CardDescription,
  CardHeader,
  CardContent,
  Card,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";

export default function RouteDetail({
  route,
}: {
  route: google.maps.DirectionsRoute;
}) {
  const legs = route.legs;
  if (!legs[0].departure_time || !legs[0].arrival_time) return null;
  const departureTime = new Date(legs[0].departure_time?.value);
  const arrivalTime = new Date(legs[0].arrival_time?.value);
  const diffTime = Math.floor(
    (arrivalTime.getTime() - departureTime.getTime()) / 1000 / 60
  );
  return (
    <Card>
      <CardHeader>
        <CardTitle>Train Information</CardTitle>
        <CardDescription>
          Get detailed information about your train journey
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-4 py-4 border-y">
        <div className="grid grid-cols-2 gap-2">
          <Label htmlFor="takeTime">Take Time</Label>
          <Badge className="col-span-1" color="blue" variant="default">
            <ClockIcon className="w-4 h-4 mr-2" />
            {diffTime} Minutes
          </Badge>
        </div>
        <div className="grid grid-cols-2 gap-2">
          <Label htmlFor="steps">Steps</Label>
          <Badge className="col-span-1" color="green" variant="default">
            <ListIcon className="w-4 h-4 mr-2" />
            {legs[0].steps.length} Steps
          </Badge>
        </div>
        <div className="grid grid-cols-2 gap-2">
          <Label htmlFor="timeToTime">Time to Time</Label>
          <Badge className="col-span-1" color="yellow" variant="default">
            <CalendarIcon className="w-4 h-4 mr-2" />
            {legs[0].departure_time?.text} - {legs[0].arrival_time?.text}
          </Badge>
        </div>
        <div className="grid grid-cols-2 gap-2">
          <Label htmlFor="oneWay">One Way</Label>
          <Badge className="col-span-1" color="red" variant="default">
            <ArrowRightIcon className="w-4 h-4 mr-2" />
            Yes
          </Badge>
        </div>
      </CardContent>
    </Card>
  );
}

function ClockIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <circle cx="12" cy="12" r="10" />
      <polyline points="12 6 12 12 16 14" />
    </svg>
  );
}

function ListIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <line x1="8" x2="21" y1="6" y2="6" />
      <line x1="8" x2="21" y1="12" y2="12" />
      <line x1="8" x2="21" y1="18" y2="18" />
      <line x1="3" x2="3.01" y1="6" y2="6" />
      <line x1="3" x2="3.01" y1="12" y2="12" />
      <line x1="3" x2="3.01" y1="18" y2="18" />
    </svg>
  );
}

function CalendarIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect width="18" height="18" x="3" y="4" rx="2" ry="2" />
      <line x1="16" x2="16" y1="2" y2="6" />
      <line x1="8" x2="8" y1="2" y2="6" />
      <line x1="3" x2="21" y1="10" y2="10" />
    </svg>
  );
}

function ArrowRightIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M5 12h14" />
      <path d="m12 5 7 7-7 7" />
    </svg>
  );
}
