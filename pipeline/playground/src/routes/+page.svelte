<script>
  import { enhance } from '$app/forms';
</script>

<form method="POST" action="?/start" class="flex flex-col gap-4" use:enhance>
  <div>
    <label>
      Sensor
      (Send)
      <select name="sensor_send">
        <option value="local">Local</option>
        <option value="cloud">Cloud</option>
      </select>
    </label>
  </div>
  <div>
    <label>
      Compute
      (Receive)
      <select name="compute_receive">
        <option value="local">Local</option>
        <option value="cloud">Cloud</option>
      </select>
      (Send)
      <select name="compute_send">
        <option value="local">Local</option>
        <option value="cloud">Cloud</option>
      </select>
    </label>
  </div>
  <div>
    <label>
      Receiver
      (Receive)
      <select name="receiver_receive">
        <option value="local">Local</option>
        <option value="cloud">Cloud</option>
      </select>
    </label>
  </div>
  <div>
    <input type="submit" value="Start" />
  </div>
</form>
