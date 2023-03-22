<script>
  import { enhance } from "$app/forms";
</script>

<form method="POST" action="?/start" class="flex flex-col gap-4" use:enhance>
  <div class="flex flex-col gap-2">
    <h4>Sensor</h4>
    <div class="flex flex-row gap-4">
      <label class="label">
        <span>Receive</span>
        <select class="select" name="sensor_receive" disabled>
          <option value="local">Local</option>
          <option value="cloud">Cloud</option>
        </select>
      </label>
      <label class="label">
        <span>Send</span>
        <select class="select" name="sensor_send">
          <option value="local">Local</option>
          <option value="cloud">Cloud</option>
        </select>
      </label>
    </div>
  </div>

  <div class="flex flex-col gap-2">
    <h4>Compute</h4>
    <div class="flex flex-row gap-4">
      <label class="label">
        <span>Receive</span>
        <select class="select" name="compute_receive">
          <option value="local">Local</option>
          <option value="cloud">Cloud</option>
        </select>
      </label>
      <label class="label">
        <span>Send</span>
        <select class="select" name="compute_send">
          <option value="local">Local</option>
          <option value="cloud">Cloud</option>
        </select>
      </label>
    </div>
  </div>

  <div class="flex flex-col gap-2">
    <h4>Receiver</h4>
    <div class="flex flex-row gap-4">
      <label class="label">
        <span>Receive</span>
        <select class="select" name="receiver_receive">
          <option value="local">Local</option>
          <option value="cloud">Cloud</option>
        </select>
      </label>
      <label class="label">
        <span>Send</span>
        <select class="select" name="receiver_send" disabled>
          <option value="local">Local</option>
          <option value="cloud">Cloud</option>
        </select>
      </label>
    </div>
  </div>

  <div>
    <input type="submit" value="Start" class="btn variant-filled" />
  </div>
</form>
