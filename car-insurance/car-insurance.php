<?php
declare(strict_types = 1);
require 'cove/data.php';
?>
<!doctype html>
<html lang="en">
<head>
   <link rel="stylesheet" href="cove.css">
   <meta charset="utf-8">
   <title>Car insurance</title>
</head>
<body>
<?php
$date_o = new DateTime;
$add_o = new DateInterval('P1M');

for ($n = 12; $n > 0; $n--) {
   $from_s = $date_o->format('Y-m-d');
   $date_o->add($add_o);
   $to_s = $date_o->format('Y-m-d');
?>
   <h1>TEXAS LIABILITY INSURANCE CARD</h1>
   <table>
      <tr>
         <td>
            <b>Name and Address of Insured<br>
Nombre y Direction del Asegurado</b>
         </td>
         <td>
            <b>Insurance Company &mdash; Compania de Seguro<br>
<?= $company1_s ?></b>
         </td>
      </tr>
      <tr>
         <td><?= $driver_name_s ?><br>
<?= $address1_s ?></td>
         <td><?= $company2_s ?><br>
<?= $company_phone_s ?></td>
      </tr>
      <tr>
         <td><?= $address2_s ?></td>
         <td>
            <span>Agent &mdash; Agente<br>
<?= $agent_name_s ?></span>
            <span>Phone #: <?= $agent_phone_s ?></span>
         </td>
      </tr>
      <tr>
         <td>Policy Number &mdash; Numero de Poliza</td>
         <td><?= $policy_num_s ?></td>
      </tr>
      <tr>
         <td>Effective Date &mdash; Fecha Efectiva</td>
         <td><?= $from_s ?></td>
      </tr>
      <tr>
         <td>Expiration Date &mdash; Fecha de Expiration</td>
         <td><?= $to_s ?> &mdash; 12:01 A.M.</td>
      </tr>
      <tr>
         <td>
            <b>Vehicle Year / Make / Model<br>
Vehiculo Ane / Marca / Modele</b>
         </td>
         <td><?= $make_model_s ?><br>
<?= $vehicle_num_s ?></td>
      </tr>
      <tr>
         <td>This policy provides at least the minimum amounts of liability
insurance required by the Texas Motor Vehicle Safety Responsibility Act for the
specified vehicle and named insured and that provided coverage for other persons
and other vehicles as provided by the insurance policy.</td>
         <td>Esta poliza provee por lo menos la cantidad minima de seguro de
responsabilidad requerida por ley (Texas Motor Vehicle Safety Responsibility
Act) para el vehiculo especificado y para los asegurados nombrados, y puede
proveer cobertura para lotras personas y otros vehiculuos segun provisto en la
poliza de seguro.</td>
      </tr>
      <tr>
         <td colspan="2">
            <b>Driver Name(s)</b>
         </td>
      </tr>
      <tr>
         <td colspan="2">
            <ol>
               <li><?= $driver_name_s ?></li>
            </ol>
         </td>
      </tr>
      <tr>
         <td>
            <p class="center">
               <b>Texas Liability<br>Insurance Card</b>
            </p>
            <p class="center">
               <b>Keep this Card</b>
            </p>
            <p><b>Important:</b> This card or a copy of your insurance policy
must be shown when you apply for or renew your:</p>
            <ul>
               <li>Motor vehichle registration</li>
               <li>Driver’s license</li>
               <li>Motor vehicle safety inspection sticker</li>
            </ul>
            <p>You also may be asked to show this card or your policy if you
have an accident or if a police officer asks to see it.</p>
            <p>All drivers in Texas must carry liability insurance on their
vehicles or otherwise meet legal requirements for financial responsibility.
Failure to do so could result in fines up to $1,000, suspension of your driver’s
license and motor vehicle registration, and impoundment of your vehicle for up
to 180 days (at a cost of $15 per day).</p>
         </td>
         <td>
            <p class="center">
               <b>Tarjeta de Seguro de<br>Responsabilidad de Texas</b>
            </p>
            <p class="center">
               <b>Guarde esta tarjeta</b>
            </p>
            <p><b>Importantre:</b> Esta tarjeta o una copia de su poliza de
seguro debe ser mostrada cuando usted solicite o renueve su:</p>
            <ul>
               <li>Registro de vehiculo de motor</li>
               <li>Licencia para conducir</li>
               <li>Etiqueta de inspeccion de seguidad para su vehiculo</li>
            </ul>
            <p>Puede que usted tenga tambien que mostrar esta tarjeta o su
poliza de sequro si tiene un accidente o si un oficial de la paz se la pide.</p>
            <p>Todos los conductors in Texas deben de tener seguro de
responsabilidad para sus vehiculos, o de otra manera llenar los requisitos
legales de responsabilidad civil. Fallo en llenar este requisito pudiera
resultar en multas de hasta $1,000, suspension de su licencia para conducir y su
registro de vehiculo de motor, y la retencion de su vehiculo por un periodo de
hasta 180 dias (a un costo de $15 por dia).</p>
         </td>
      </tr>
   </table>
<?php
}
?>
</body>
</html>
